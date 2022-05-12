package governance

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/accounts/abi/bind/backends"
	"github.com/klaytn/klaytn/blockchain"
	"github.com/klaytn/klaytn/common"
	govcontract "github.com/klaytn/klaytn/contracts/gov"
	"github.com/klaytn/klaytn/crypto"
	"github.com/klaytn/klaytn/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func populateContract(t *testing.T) ([]*bind.TransactOpts, *backends.SimulatedBackend, common.Address, *govcontract.GovParam) {
	// Create accounts and simulated blockchain
	accounts := []*bind.TransactOpts{}
	alloc := blockchain.GenesisAlloc{}
	for i := 0; i < 1; i++ {
		key, _ := crypto.GenerateKey()
		account := bind.NewKeyedTransactor(key)
		accounts = append(accounts, account)
		alloc[account.From] = blockchain.GenesisAccount{Balance: big.NewInt(params.KLAY)}
	}
	sim := backends.NewSimulatedBackend(alloc)

	// Deploy contract
	owner := accounts[0]
	address, _, contract, err := govcontract.DeployGovParam(owner, sim, owner.From)
	require.Nil(t, err)
	sim.Commit()

	// Populate initial chainconfig
	_, err = contract.AddParam(owner, big.NewInt(0), "governance.unitprice", []byte{0x12, 0x34})
	require.Nil(t, err)
	sim.Commit()

	return accounts, sim, address, contract
}

func TestContract_Call(t *testing.T) {
	_, sim, _, contract := populateContract(t)

	num := sim.BlockChain().CurrentBlock().Number()

	paramViews, err := contract.GetAllParams(&bind.CallOpts{BlockNumber: num}, big.NewInt(2))
	assert.Nil(t, err)
	fmt.Println(paramViews)
}
