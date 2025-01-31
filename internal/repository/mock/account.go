// Code generated by MockGen. DO NOT EDIT.
// Source: ./account.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/tonghia/go-challenge-transaction-app/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountRepositorier is a mock of AccountRepositorier interface.
type MockAccountRepositorier struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositorierMockRecorder
}

// MockAccountRepositorierMockRecorder is the mock recorder for MockAccountRepositorier.
type MockAccountRepositorierMockRecorder struct {
	mock *MockAccountRepositorier
}

// NewMockAccountRepositorier creates a new mock instance.
func NewMockAccountRepositorier(ctrl *gomock.Controller) *MockAccountRepositorier {
	mock := &MockAccountRepositorier{ctrl: ctrl}
	mock.recorder = &MockAccountRepositorierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepositorier) EXPECT() *MockAccountRepositorierMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockAccountRepositorier) GetByID(ctx context.Context, id int64) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAccountRepositorierMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAccountRepositorier)(nil).GetByID), ctx, id)
}
