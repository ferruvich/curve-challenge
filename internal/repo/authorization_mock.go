// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ferruvich/curve-prepaid-card/internal/repo (interfaces: AuthorizationRequest)

// Package repo is a generated GoMock package.
package repo

import (
	context "context"
	model "github.com/ferruvich/curve-prepaid-card/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthorizationRequest is a mock of AuthorizationRequest interface
type MockAuthorizationRequest struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationRequestMockRecorder
}

// MockAuthorizationRequestMockRecorder is the mock recorder for MockAuthorizationRequest
type MockAuthorizationRequestMockRecorder struct {
	mock *MockAuthorizationRequest
}

// NewMockAuthorizationRequest creates a new mock instance
func NewMockAuthorizationRequest(ctrl *gomock.Controller) *MockAuthorizationRequest {
	mock := &MockAuthorizationRequest{ctrl: ctrl}
	mock.recorder = &MockAuthorizationRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizationRequest) EXPECT() *MockAuthorizationRequestMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockAuthorizationRequest) Write(arg0 context.Context, arg1 *model.AuthorizationRequest) error {
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockAuthorizationRequestMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockAuthorizationRequest)(nil).Write), arg0, arg1)
}
