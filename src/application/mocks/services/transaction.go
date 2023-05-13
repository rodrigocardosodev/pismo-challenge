// Code generated by MockGen. DO NOT EDIT.
// Source: src/application/services/transaction.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/rodrigocardosodev/pismo-challenge/src/application/models"
)

// MockTransactionServiceInterface is a mock of TransactionServiceInterface interface.
type MockTransactionServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceInterfaceMockRecorder
}

// MockTransactionServiceInterfaceMockRecorder is the mock recorder for MockTransactionServiceInterface.
type MockTransactionServiceInterfaceMockRecorder struct {
	mock *MockTransactionServiceInterface
}

// NewMockTransactionServiceInterface creates a new mock instance.
func NewMockTransactionServiceInterface(ctrl *gomock.Controller) *MockTransactionServiceInterface {
	mock := &MockTransactionServiceInterface{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionServiceInterface) EXPECT() *MockTransactionServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransactionServiceInterface) Create(ctx context.Context, accountId int64, operationId int8, amount uint64) (models.TransactionInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, accountId, operationId, amount)
	ret0, _ := ret[0].(models.TransactionInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTransactionServiceInterfaceMockRecorder) Create(ctx, accountId, operationId, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionServiceInterface)(nil).Create), ctx, accountId, operationId, amount)
}