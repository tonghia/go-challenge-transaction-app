// Code generated by MockGen. DO NOT EDIT.
// Source: ./account_transaction.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/tonghia/go-challenge-transaction-app/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountTransactionRepositorier is a mock of AccountTransactionRepositorier interface.
type MockAccountTransactionRepositorier struct {
	ctrl     *gomock.Controller
	recorder *MockAccountTransactionRepositorierMockRecorder
}

// MockAccountTransactionRepositorierMockRecorder is the mock recorder for MockAccountTransactionRepositorier.
type MockAccountTransactionRepositorierMockRecorder struct {
	mock *MockAccountTransactionRepositorier
}

// NewMockAccountTransactionRepositorier creates a new mock instance.
func NewMockAccountTransactionRepositorier(ctrl *gomock.Controller) *MockAccountTransactionRepositorier {
	mock := &MockAccountTransactionRepositorier{ctrl: ctrl}
	mock.recorder = &MockAccountTransactionRepositorierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountTransactionRepositorier) EXPECT() *MockAccountTransactionRepositorierMockRecorder {
	return m.recorder
}

// CreateOne mocks base method.
func (m *MockAccountTransactionRepositorier) CreateOne(ctx context.Context, txn *model.AccountTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOne", ctx, txn)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOne indicates an expected call of CreateOne.
func (mr *MockAccountTransactionRepositorierMockRecorder) CreateOne(ctx, txn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOne", reflect.TypeOf((*MockAccountTransactionRepositorier)(nil).CreateOne), ctx, txn)
}

// DeleteByTransactionID mocks base method.
func (m *MockAccountTransactionRepositorier) DeleteByTransactionID(ctx context.Context, txnID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByTransactionID", ctx, txnID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByTransactionID indicates an expected call of DeleteByTransactionID.
func (mr *MockAccountTransactionRepositorierMockRecorder) DeleteByTransactionID(ctx, txnID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByTransactionID", reflect.TypeOf((*MockAccountTransactionRepositorier)(nil).DeleteByTransactionID), ctx, txnID)
}

// GetByID mocks base method.
func (m *MockAccountTransactionRepositorier) GetByID(ctx context.Context, id int64) (*model.AccountTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*model.AccountTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAccountTransactionRepositorierMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAccountTransactionRepositorier)(nil).GetByID), ctx, id)
}

// GetByUser mocks base method.
func (m *MockAccountTransactionRepositorier) GetByUser(ctx context.Context, userID int64) ([]*model.AccountTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUser", ctx, userID)
	ret0, _ := ret[0].([]*model.AccountTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUser indicates an expected call of GetByUser.
func (mr *MockAccountTransactionRepositorierMockRecorder) GetByUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUser", reflect.TypeOf((*MockAccountTransactionRepositorier)(nil).GetByUser), ctx, userID)
}

// GetByUserAccount mocks base method.
func (m *MockAccountTransactionRepositorier) GetByUserAccount(ctx context.Context, userID, accountID int64) ([]*model.AccountTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserAccount", ctx, userID, accountID)
	ret0, _ := ret[0].([]*model.AccountTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserAccount indicates an expected call of GetByUserAccount.
func (mr *MockAccountTransactionRepositorierMockRecorder) GetByUserAccount(ctx, userID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserAccount", reflect.TypeOf((*MockAccountTransactionRepositorier)(nil).GetByUserAccount), ctx, userID, accountID)
}

// UpdateOne mocks base method.
func (m *MockAccountTransactionRepositorier) UpdateOne(ctx context.Context, txn *model.AccountTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", ctx, txn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockAccountTransactionRepositorierMockRecorder) UpdateOne(ctx, txn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockAccountTransactionRepositorier)(nil).UpdateOne), ctx, txn)
}
