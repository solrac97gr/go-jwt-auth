// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports (interfaces: MiddlewareApplication)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMiddlewareApplication is a mock of MiddlewareApplication interface.
type MockMiddlewareApplication struct {
	ctrl     *gomock.Controller
	recorder *MockMiddlewareApplicationMockRecorder
}

// MockMiddlewareApplicationMockRecorder is the mock recorder for MockMiddlewareApplication.
type MockMiddlewareApplicationMockRecorder struct {
	mock *MockMiddlewareApplication
}

// NewMockMiddlewareApplication creates a new mock instance.
func NewMockMiddlewareApplication(ctrl *gomock.Controller) *MockMiddlewareApplication {
	mock := &MockMiddlewareApplication{ctrl: ctrl}
	mock.recorder = &MockMiddlewareApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddlewareApplication) EXPECT() *MockMiddlewareApplicationMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockMiddlewareApplication) Authenticate(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockMiddlewareApplicationMockRecorder) Authenticate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockMiddlewareApplication)(nil).Authenticate), arg0)
}
