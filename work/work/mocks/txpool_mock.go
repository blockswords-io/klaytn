// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/klaytn/klaytn/work (interfaces: TxPool)

// Package mocks is a generated GoMock package.
package mocks

import (
	big "math/big"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	blockchain "github.com/klaytn/klaytn/blockchain"
	types "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// MockTxPool is a mock of TxPool interface
type MockTxPool struct {
	ctrl     *gomock.Controller
	recorder *MockTxPoolMockRecorder
}

// MockTxPoolMockRecorder is the mock recorder for MockTxPool
type MockTxPoolMockRecorder struct {
	mock *MockTxPool
}

// NewMockTxPool creates a new mock instance
func NewMockTxPool(ctrl *gomock.Controller) *MockTxPool {
	mock := &MockTxPool{ctrl: ctrl}
	mock.recorder = &MockTxPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxPool) EXPECT() *MockTxPoolMockRecorder {
	return m.recorder
}

// AddLocal mocks base method
func (m *MockTxPool) AddLocal(arg0 *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLocal", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLocal indicates an expected call of AddLocal
func (mr *MockTxPoolMockRecorder) AddLocal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLocal", reflect.TypeOf((*MockTxPool)(nil).AddLocal), arg0)
}

// CachedPendingTxsByCount mocks base method
func (m *MockTxPool) CachedPendingTxsByCount(arg0 int) types.Transactions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CachedPendingTxsByCount", arg0)
	ret0, _ := ret[0].(types.Transactions)
	return ret0
}

// CachedPendingTxsByCount indicates an expected call of CachedPendingTxsByCount
func (mr *MockTxPoolMockRecorder) CachedPendingTxsByCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CachedPendingTxsByCount", reflect.TypeOf((*MockTxPool)(nil).CachedPendingTxsByCount), arg0)
}

// Content mocks base method
func (m *MockTxPool) Content() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Content")
	ret0, _ := ret[0].(map[common.Address]types.Transactions)
	ret1, _ := ret[1].(map[common.Address]types.Transactions)
	return ret0, ret1
}

// Content indicates an expected call of Content
func (mr *MockTxPoolMockRecorder) Content() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Content", reflect.TypeOf((*MockTxPool)(nil).Content))
}

// GasPrice mocks base method
func (m *MockTxPool) GasPrice() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasPrice")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GasPrice indicates an expected call of GasPrice
func (mr *MockTxPoolMockRecorder) GasPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasPrice", reflect.TypeOf((*MockTxPool)(nil).GasPrice))
}

// Get mocks base method
func (m *MockTxPool) Get(arg0 common.Hash) *types.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*types.Transaction)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockTxPoolMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTxPool)(nil).Get), arg0)
}

// GetPendingNonce mocks base method
func (m *MockTxPool) GetPendingNonce(arg0 common.Address) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingNonce", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetPendingNonce indicates an expected call of GetPendingNonce
func (mr *MockTxPoolMockRecorder) GetPendingNonce(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingNonce", reflect.TypeOf((*MockTxPool)(nil).GetPendingNonce), arg0)
}

// HandleTxMsg mocks base method
func (m *MockTxPool) HandleTxMsg(arg0 types.Transactions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleTxMsg", arg0)
}

// HandleTxMsg indicates an expected call of HandleTxMsg
func (mr *MockTxPoolMockRecorder) HandleTxMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTxMsg", reflect.TypeOf((*MockTxPool)(nil).HandleTxMsg), arg0)
}

// Pending mocks base method
func (m *MockTxPool) Pending() (map[common.Address]types.Transactions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pending")
	ret0, _ := ret[0].(map[common.Address]types.Transactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pending indicates an expected call of Pending
func (mr *MockTxPoolMockRecorder) Pending() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pending", reflect.TypeOf((*MockTxPool)(nil).Pending))
}

// SetGasPrice mocks base method
func (m *MockTxPool) SetGasPrice(arg0 *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGasPrice", arg0)
}

// SetGasPrice indicates an expected call of SetGasPrice
func (mr *MockTxPoolMockRecorder) SetGasPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGasPrice", reflect.TypeOf((*MockTxPool)(nil).SetGasPrice), arg0)
}

// StartSpamThrottler mocks base method
func (m *MockTxPool) StartSpamThrottler(arg0 *blockchain.ThrottlerConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSpamThrottler", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartSpamThrottler indicates an expected call of StartSpamThrottler
func (mr *MockTxPoolMockRecorder) StartSpamThrottler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSpamThrottler", reflect.TypeOf((*MockTxPool)(nil).StartSpamThrottler), arg0)
}

// Stats mocks base method
func (m *MockTxPool) Stats() (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockTxPoolMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockTxPool)(nil).Stats))
}

// Stop mocks base method
func (m *MockTxPool) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockTxPoolMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTxPool)(nil).Stop))
}

// StopSpamThrottler mocks base method
func (m *MockTxPool) StopSpamThrottler() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopSpamThrottler")
}

// StopSpamThrottler indicates an expected call of StopSpamThrottler
func (mr *MockTxPoolMockRecorder) StopSpamThrottler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopSpamThrottler", reflect.TypeOf((*MockTxPool)(nil).StopSpamThrottler))
}

// SubscribeNewTxsEvent mocks base method
func (m *MockTxPool) SubscribeNewTxsEvent(arg0 chan<- blockchain.NewTxsEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewTxsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeNewTxsEvent indicates an expected call of SubscribeNewTxsEvent
func (mr *MockTxPoolMockRecorder) SubscribeNewTxsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewTxsEvent", reflect.TypeOf((*MockTxPool)(nil).SubscribeNewTxsEvent), arg0)
}
