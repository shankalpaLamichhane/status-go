// Code generated by MockGen. DO NOT EDIT.
// Source: transactions/fake/txservice.go

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	common "github.com/ethereum/go-ethereum/common"
	hexutil "github.com/ethereum/go-ethereum/common/hexutil"
	rpc "github.com/ethereum/go-ethereum/rpc"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPublicTransactionPoolAPI is a mock of PublicTransactionPoolAPI interface
type MockPublicTransactionPoolAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPublicTransactionPoolAPIMockRecorder
}

// MockPublicTransactionPoolAPIMockRecorder is the mock recorder for MockPublicTransactionPoolAPI
type MockPublicTransactionPoolAPIMockRecorder struct {
	mock *MockPublicTransactionPoolAPI
}

// NewMockPublicTransactionPoolAPI creates a new mock instance
func NewMockPublicTransactionPoolAPI(ctrl *gomock.Controller) *MockPublicTransactionPoolAPI {
	mock := &MockPublicTransactionPoolAPI{ctrl: ctrl}
	mock.recorder = &MockPublicTransactionPoolAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublicTransactionPoolAPI) EXPECT() *MockPublicTransactionPoolAPIMockRecorder {
	return m.recorder
}

// GasPrice mocks base method
func (m *MockPublicTransactionPoolAPI) GasPrice(ctx context.Context) (*hexutil.Big, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasPrice", ctx)
	ret0, _ := ret[0].(*hexutil.Big)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasPrice indicates an expected call of GasPrice
func (mr *MockPublicTransactionPoolAPIMockRecorder) GasPrice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasPrice", reflect.TypeOf((*MockPublicTransactionPoolAPI)(nil).GasPrice), ctx)
}

// EstimateGas mocks base method
func (m *MockPublicTransactionPoolAPI) EstimateGas(ctx context.Context, args CallArgs) (hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", ctx, args)
	ret0, _ := ret[0].(hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas
func (mr *MockPublicTransactionPoolAPIMockRecorder) EstimateGas(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockPublicTransactionPoolAPI)(nil).EstimateGas), ctx, args)
}

// GetTransactionCount mocks base method
func (m *MockPublicTransactionPoolAPI) GetTransactionCount(ctx context.Context, address common.Address, blockNr rpc.BlockNumber) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionCount", ctx, address, blockNr)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionCount indicates an expected call of GetTransactionCount
func (mr *MockPublicTransactionPoolAPIMockRecorder) GetTransactionCount(ctx, address, blockNr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCount", reflect.TypeOf((*MockPublicTransactionPoolAPI)(nil).GetTransactionCount), ctx, address, blockNr)
}

// SendRawTransaction mocks base method
func (m *MockPublicTransactionPoolAPI) SendRawTransaction(ctx context.Context, encodedTx hexutil.Bytes) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRawTransaction", ctx, encodedTx)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTransaction indicates an expected call of SendRawTransaction
func (mr *MockPublicTransactionPoolAPIMockRecorder) SendRawTransaction(ctx, encodedTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTransaction", reflect.TypeOf((*MockPublicTransactionPoolAPI)(nil).SendRawTransaction), ctx, encodedTx)
}
