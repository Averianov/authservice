// Code generated by MockGen. DO NOT EDIT.
// Source: authservice/models (interfaces: Session)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSession is a mock of Session interface
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// CompareWithExisting mocks base method
func (m *MockSession) CompareWithExisting(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareWithExisting", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompareWithExisting indicates an expected call of CompareWithExisting
func (mr *MockSessionMockRecorder) CompareWithExisting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareWithExisting", reflect.TypeOf((*MockSession)(nil).CompareWithExisting), arg0)
}

// Save mocks base method
func (m *MockSession) Save() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockSessionMockRecorder) Save() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockSession)(nil).Save))
}
