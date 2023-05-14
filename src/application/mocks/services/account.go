// Code generated by MockGen. DO NOT EDIT.
// Source: src/application/services/account.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/rodrigocardosodev/pismo-challenge/src/application/models"
)

// MockIAccountService is a mock of IAccountService interface.
type MockIAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountServiceMockRecorder
}

// MockIAccountServiceMockRecorder is the mock recorder for MockIAccountService.
type MockIAccountServiceMockRecorder struct {
	mock *MockIAccountService
}

// NewMockIAccountService creates a new mock instance.
func NewMockIAccountService(ctrl *gomock.Controller) *MockIAccountService {
	mock := &MockIAccountService{ctrl: ctrl}
	mock.recorder = &MockIAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountService) EXPECT() *MockIAccountServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAccountService) Create(ctx context.Context, documentNumber string) (models.AccountInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, documentNumber)
	ret0, _ := ret[0].(models.AccountInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIAccountServiceMockRecorder) Create(ctx, documentNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAccountService)(nil).Create), ctx, documentNumber)
}

// GetByID mocks base method.
func (m *MockIAccountService) GetByID(ctx context.Context, id int64) (models.AccountInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(models.AccountInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIAccountServiceMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIAccountService)(nil).GetByID), ctx, id)
}

// IsValidCPF mocks base method.
func (m *MockIAccountService) IsValidCPF(cpf string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidCPF", cpf)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsValidCPF indicates an expected call of IsValidCPF.
func (mr *MockIAccountServiceMockRecorder) IsValidCPF(cpf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidCPF", reflect.TypeOf((*MockIAccountService)(nil).IsValidCPF), cpf)
}
