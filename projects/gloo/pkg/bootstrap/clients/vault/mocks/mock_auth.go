// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hashicorp/vault/api (interfaces: AuthMethod)
//
// Generated by this command:
//
//	mockgen -destination mocks/mock_auth.go -package mocks github.com/hashicorp/vault/api AuthMethod
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	api "github.com/hashicorp/vault/api"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthMethod is a mock of AuthMethod interface.
type MockAuthMethod struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMethodMockRecorder
}

// MockAuthMethodMockRecorder is the mock recorder for MockAuthMethod.
type MockAuthMethodMockRecorder struct {
	mock *MockAuthMethod
}

// NewMockAuthMethod creates a new mock instance.
func NewMockAuthMethod(ctrl *gomock.Controller) *MockAuthMethod {
	mock := &MockAuthMethod{ctrl: ctrl}
	mock.recorder = &MockAuthMethodMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthMethod) EXPECT() *MockAuthMethodMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthMethod) Login(arg0 context.Context, arg1 *api.Client) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMethodMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthMethod)(nil).Login), arg0, arg1)
}