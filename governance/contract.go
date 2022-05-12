package governance

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/blockchain"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/blockchain/vm"
	"github.com/klaytn/klaytn/common"
	govcontract "github.com/klaytn/klaytn/contracts/gov"
	"github.com/klaytn/klaytn/params"
)

type ContractEngine struct {
	initialConfig *params.ChainConfig
	initialParams *params.GovParamSet
	currentParams *params.GovParamSet

	chain   blockChain
	address common.Address // GovParams contract address
	abi     abi.ABI        // parsed contract ABI
}

func NewContractEngine(config *params.ChainConfig) *ContractEngine {
	e := &ContractEngine{
		address: config.Governance.GovParamsContract,
	}

	if govParams, err := params.NewGovParamSetChainConfig(config); err == nil {
		e.initialConfig = config
		e.initialParams = govParams
		e.currentParams = govParams
	} else {
		logger.Crit("Error parsing initial ChainConfig", "err", err)
	}

	if abi, err := abi.JSON(strings.NewReader(govcontract.GovParamABI)); err != nil {
		logger.Crit("Cannot parse GovParamABI", "err", err)
	} else {
		e.abi = abi
	}

	return e
}

// Params effective at upcoming block (head+1), queried current block (head)
func (e *ContractEngine) Params() *params.GovParamSet {
	return e.currentParams
}

// Params effective at requested block (num), queried at previous block (num-1).
func (e *ContractEngine) ParamsAt(num uint64) (*params.GovParamSet, error) {
	if num == 0 {
		num = 1
	}
	stateBlock := new(big.Int).SetUint64(num - 1)
	paramBlock := new(big.Int).SetUint64(num)
	return e.contractGetAllParams(stateBlock, paramBlock)
}

func (e *ContractEngine) UpdateParams() error {
	head := e.chain.CurrentHeader().Number.Uint64()
	govParams, err := e.ParamsAt(head + 1)
	if err != nil {
		return err
	}
	e.currentParams = govParams
	return nil
}

func (e *ContractEngine) SetBlockchain(chain blockChain) {
	e.chain = chain
}

// At block `stateBlock`, call function `getAllParams(paramBlock)`.
func (e *ContractEngine) contractGetAllParams(stateBlock, paramBlock *big.Int) (*params.GovParamSet, error) {
	tx, err := e.makeTx(stateBlock, "getAllParams", paramBlock)
	if err != nil {
		return nil, err
	}

	res, err := e.callTx(stateBlock, tx)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return params.NewGovParamSet(), nil
	}

	var paramViews []govcontract.GovParamParamView
	if err := e.abi.Unpack(&paramViews, "getAllParams", res); err != nil {
		return nil, err
	}
	bytesMap := make(map[string][]byte)
	for _, pair := range paramViews {
		bytesMap[pair.Name] = pair.Value
	}

	return params.NewGovParamSetBytesMap(bytesMap)
}

// Make contract execution transaction for calling `fn` function with `args` at block number `num`.
func (e *ContractEngine) makeTx(num *big.Int, fn string, args ...interface{}) (*types.Transaction, error) {
	calldata, err := e.abi.Pack(fn, args...)
	if err != nil {
		return nil, err
	}

	rules := params.CypressChainConfig.Rules(num)
	intrinsicGas, err := types.IntrinsicGas(calldata, nil, false, rules)
	if err != nil {
		return nil, err
	}

	var (
		from       = common.Address{}
		to         = &e.address
		nonce      = uint64(0)
		amount     = big.NewInt(0)
		gasLimit   = uint64(1e8)
		gasPrice   = big.NewInt(0)
		checkNonce = false
	)
	tx := types.NewMessage(from, to, nonce, amount, gasLimit, gasPrice,
		calldata, checkNonce, intrinsicGas)
	return tx, nil
}

func (e *ContractEngine) callTx(num *big.Int, tx *types.Transaction) ([]byte, error) {
	if e.chain == nil {
		return nil, errors.New("blockchain not initiatorEncHandshakeialized")
	}

	// Load state at given block
	block := e.chain.GetBlockByNumber(num.Uint64())
	if block == nil {
		return nil, fmt.Errorf("Cannot get block %d", num.Uint64())
	}
	statedb, err := e.chain.StateAt(block.Root())
	if err != nil {
		return nil, err
	}

	// Run EVM at given states
	evmCtx := blockchain.NewEVMContext(tx, block.Header(), e.chain, nil)
	evm := vm.NewEVM(evmCtx, statedb, e.chain.Config(), &vm.Config{})

	res, _, kerr := blockchain.ApplyMessage(evm, tx)
	if kerr.ErrTxInvalid != nil {
		return nil, kerr.ErrTxInvalid
	}

	return res, nil
}
