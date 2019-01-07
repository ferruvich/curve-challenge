// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ferruvich/curve-challenge/internal/repo (interfaces: Card)

// Package repo is a generated GoMock package.
package repo

import (
	context "context"
	reflect "reflect"

	model "github.com/ferruvich/curve-challenge/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockCard is a mock of Card interface
type MockCard struct {
	ctrl     *gomock.Controller
	recorder *MockCardMockRecorder
}

// MockCardMockRecorder is the mock recorder for MockCard
type MockCardMockRecorder struct {
	mock *MockCard
}

// NewMockCard creates a new mock instance
func NewMockCard(ctrl *gomock.Controller) *MockCard {
	mock := &MockCard{ctrl: ctrl}
	mock.recorder = &MockCardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCard) EXPECT() *MockCardMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockCard) Write(arg0 context.Context, arg1 *model.Card) error {
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockCardMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCard)(nil).Write), arg0, arg1)
}
