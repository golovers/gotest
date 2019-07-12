// Code generated by MockGen. DO NOT EDIT.
// Source: user_handler.go

// Package user_test is a generated GoMock package.
package user_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	user "github.com/golovers/gotest/user"
	reflect "reflect"
)

// Mockservice is a mock of service interface
type Mockservice struct {
	ctrl     *gomock.Controller
	recorder *MockserviceMockRecorder
}

// MockserviceMockRecorder is the mock recorder for Mockservice
type MockserviceMockRecorder struct {
	mock *Mockservice
}

// NewMockservice creates a new mock instance
func NewMockservice(ctrl *gomock.Controller) *Mockservice {
	mock := &Mockservice{ctrl: ctrl}
	mock.recorder = &MockserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockservice) EXPECT() *MockserviceMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *Mockservice) Register(ctx context.Context, user *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockserviceMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*Mockservice)(nil).Register), ctx, user)
}