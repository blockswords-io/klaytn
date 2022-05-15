package api

import (
	"context"
	"math/big"
	"testing"

	"github.com/golang/mock/gomock"
	mock_api "github.com/klaytn/klaytn/api/mocks"
	"github.com/klaytn/klaytn/blockchain"
	"github.com/klaytn/klaytn/blockchain/state"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/storage/database"
	"github.com/stretchr/testify/assert"
)

func testCreateNewPublicBlockChainAPI(t *testing.T) (*mock_api.MockBackend, *PublicBlockChainAPI) {
	mockCtrl := gomock.NewController(t)
	mockBackend := mock_api.NewMockBackend(mockCtrl)
	api := NewPublicBlockChainAPI(mockBackend)

	blockchain.InitDeriveSha(types.ImplDeriveShaOriginal)

	return mockBackend, api
}

func TestEstimateGasWithTrace(t *testing.T) {
	mb, a := testCreateNewPublicBlockChainAPI(t)

	memDBManager := database.NewMemoryDBManager()
	sdb := state.NewDatabase(memDBManager)

	mb.EXPECT().RPCGasCap().Return(nil)
	// (*state.StateDB, *types.Header, error)
	mb.EXPECT().StateAndHeaderByNumberOrHash(gomock.Any(), gomock.Any()).Return(sdb)
	res, err := a.EstimateGasWithTrace(context.Background(), CallArgs{
		From:  common.HexToAddress("0x2eaad2bf70a070aaa2e007beee99c6148f47718e"),
		To:    nil,
		Value: hexutil.Big(*new(big.Int).SetUint64(100)),
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
