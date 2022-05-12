package tests

import (
	"encoding/json"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/blockchain"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/compiler"
	"github.com/klaytn/klaytn/governance"
	"github.com/klaytn/klaytn/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func prepareGovParamContract(t *testing.T) (compiler.Contract, abi.ABI) {
	var contract compiler.Contract
	contracts, err := compiler.CompileSolidity("", "../contracts/gov/GovParam.sol")
	require.Nil(t, err)
	for k, v := range contracts {
		if strings.Contains(k, "GovParam.sol") {
			contract = *v
			break
		}
	}
	require.NotNil(t, contract)

	abiJson, err := json.Marshal(contract.Info.AbiDefinition)
	require.Nil(t, err)
	contractAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	require.Nil(t, err)

	return contract, contractAbi
}

func waitReceipt(chain *blockchain.BlockChain, txhash common.Hash) *types.Receipt {
	for i := 0; i < 30; i++ {
		time.Sleep(1 * time.Second)
		receipt := chain.GetReceiptByTxHash(txhash)
		if receipt != nil {
			return receipt
		}
	}
	return nil
}

func TestGovernanceContract(t *testing.T) {
	// enableLog()

	fullNode, node, validator, chainId, workspace := newBlockchain(t)
	defer os.RemoveAll(workspace)
	defer fullNode.Stop()

	owner := validator
	contract, contractAbi := prepareGovParamContract(t)

	var (
		signer   = types.LatestSignerForChainID(chainId)
		chain    = node.BlockChain().(*blockchain.BlockChain)
		amount   = new(big.Int).SetUint64(0)
		gasPrice = new(big.Int).SetUint64(node.BlockChain().Config().UnitPrice)
		gasLimit = uint64(100e9)

		contractAddr    = common.Address{}
		deployBlock     = uint64(0)
		addParamBlock   = uint64(0)
		setParamBlock   = uint64(0)
		activationBlock = uint64(0)

		paramId     = big.NewInt(0)
		paramName   = "governance.unitprice"
		paramValueA = uint64(25e9)
		paramValueB = uint64(750e9)
		paramBytesA = new(big.Int).SetUint64(paramValueA).Bytes()
		paramBytesB = new(big.Int).SetUint64(paramValueB).Bytes()
	)

	{ // Deploy contract
		// constructor(address _owner)
		ctorArgs, err := contractAbi.Pack("", owner.Addr)
		require.Nil(t, err)
		code := append(common.FromHex(contract.Code), ctorArgs...)

		tx, err := types.NewTransactionWithMap(types.TxTypeSmartContractDeploy, map[types.TxValueKeyType]interface{}{
			types.TxValueKeyNonce:         owner.Nonce,
			types.TxValueKeyFrom:          owner.Addr,
			types.TxValueKeyTo:            (*common.Address)(nil),
			types.TxValueKeyAmount:        amount,
			types.TxValueKeyGasLimit:      gasLimit,
			types.TxValueKeyGasPrice:      gasPrice,
			types.TxValueKeyData:          code,
			types.TxValueKeyHumanReadable: false,
			types.TxValueKeyCodeFormat:    params.CodeFormatEVM,
		})
		require.Nil(t, err)
		require.Nil(t, tx.SignWithKeys(signer, owner.Keys))
		require.Nil(t, node.TxPool().AddLocal(tx))

		receipt := waitReceipt(chain, tx.Hash())
		require.NotNil(t, receipt)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

		owner.Nonce += 1
		contractAddr = receipt.ContractAddress
		_, _, deployBlock, _ = chain.GetTxAndLookupInfo(tx.Hash())
		t.Logf("GovParam deployed at block=%d addr=%s", deployBlock, contractAddr.Hex())
	}

	{ // Add parameter
		// addParam(uint id, string memory _name, bytes memory value)
		data, err := contractAbi.Pack("addParam", paramId, paramName, paramBytesA)
		require.Nil(t, err)

		tx, err := types.NewTransactionWithMap(types.TxTypeSmartContractExecution, map[types.TxValueKeyType]interface{}{
			types.TxValueKeyNonce:    owner.Nonce,
			types.TxValueKeyFrom:     owner.Addr,
			types.TxValueKeyTo:       contractAddr,
			types.TxValueKeyAmount:   amount,
			types.TxValueKeyGasLimit: gasLimit,
			types.TxValueKeyGasPrice: gasPrice,
			types.TxValueKeyData:     data,
		})
		require.Nil(t, err)
		require.Nil(t, tx.SignWithKeys(signer, owner.Keys))
		require.Nil(t, node.TxPool().AddLocal(tx))

		receipt := waitReceipt(chain, tx.Hash())
		require.NotNil(t, receipt)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

		owner.Nonce += 1
		_, _, addParamBlock, _ = chain.GetTxAndLookupInfo(tx.Hash())
		t.Logf("GovParam.addParam success at block=%d", addParamBlock)
	}

	{ // Set parameter
		// setParam(uint id, bytes calldata value, uint64 _activationBlock)
		activationBlock = chain.CurrentHeader().Number.Uint64() + 10
		data, err := contractAbi.Pack("setParam", paramId, paramBytesB, activationBlock)
		require.Nil(t, err)

		tx, err := types.NewTransactionWithMap(types.TxTypeSmartContractExecution, map[types.TxValueKeyType]interface{}{
			types.TxValueKeyNonce:    owner.Nonce,
			types.TxValueKeyFrom:     owner.Addr,
			types.TxValueKeyTo:       contractAddr,
			types.TxValueKeyAmount:   amount,
			types.TxValueKeyGasLimit: gasLimit,
			types.TxValueKeyGasPrice: gasPrice,
			types.TxValueKeyData:     data,
		})
		require.Nil(t, err)
		require.Nil(t, tx.SignWithKeys(signer, owner.Keys))
		require.Nil(t, node.TxPool().AddLocal(tx))

		receipt := waitReceipt(chain, tx.Hash())
		require.NotNil(t, receipt)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

		owner.Nonce += 1
		_, _, setParamBlock, _ = chain.GetTxAndLookupInfo(tx.Hash())
		t.Logf("GovParam.setParam success at block=%d activation=%d", setParamBlock, activationBlock)
	}

	config := chain.Config()
	config.Governance.GovParamsContract = contractAddr
	engine := governance.NewContractEngine(config)
	engine.SetBlockchain(chain)
	_ = engine

	// Wait until activation block and verify current Params() at each block.
	chainEventCh := make(chan blockchain.ChainEvent)
	subscription := chain.SubscribeChainEvent(chainEventCh)
	defer subscription.Unsubscribe()
	timeout := time.NewTimer(15 * time.Second)
chainEventLoop:
	for {
		select {
		case <-timeout.C:
			t.Fatal("timed out")
		case ev := <-chainEventCh:
			head := ev.Block.Number().Uint64()
			err := engine.UpdateParams()
			assert.Nil(t, err)

			// Params() holds the parameters effective at (head+1).
			value := engine.Params().UnitPrice()
			t.Logf("Params()[%s] after block %d is %d", paramName, head, value)
			if head+1 < activationBlock {
				assert.Equal(t, paramValueA, value)
			} else if head+1 >= activationBlock {
				assert.Equal(t, paramValueB, value)
			}

			if head > activationBlock+1 {
				break chainEventLoop
			}
		}
	}

	// Verify historic parameters
	for num := uint64(0); num <= activationBlock+2; num++ {
		// ParamsAt() does not fail even before contract is deployed;
		// it simply returns an empty result.
		p, err := engine.ParamsAt(num)
		assert.Nil(t, err)
		require.NotNil(t, p)

		v, ok := p.Get(params.UnitPrice)
		if num <= addParamBlock {
			assert.False(t, ok) // param not yet exists before addParam
			continue
		} else {
			assert.True(t, ok)
		}

		value := v.(uint64)
		t.Logf("ParamsAt(%d)[%s] is %d", num, paramName, value)
		if num < activationBlock {
			assert.Equal(t, paramValueA, value)
		} else {
			assert.Equal(t, paramValueB, value)
		}
	}
}
