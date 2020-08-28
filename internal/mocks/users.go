// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: UsersDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsersDescriber is a mock of UsersDescriber interface
type MockUsersDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockUsersDescriberMockRecorder
}

// MockUsersDescriberMockRecorder is the mock recorder for MockUsersDescriber
type MockUsersDescriberMockRecorder struct {
	mock *MockUsersDescriber
}

// NewMockUsersDescriber creates a new mock instance
func NewMockUsersDescriber(ctrl *gomock.Controller) *MockUsersDescriber {
	mock := &MockUsersDescriber{ctrl: ctrl}
	mock.recorder = &MockUsersDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersDescriber) EXPECT() *MockUsersDescriberMockRecorder {
	return m.recorder
}

// UserByID mocks base method
func (m *MockUsersDescriber) UserByID(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByID", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserByID indicates an expected call of UserByID
func (mr *MockUsersDescriberMockRecorder) UserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByID", reflect.TypeOf((*MockUsersDescriber)(nil).UserByID), arg0)
}

// UserByName mocks base method
func (m *MockUsersDescriber) UserByName(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByName", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserByName indicates an expected call of UserByName
func (mr *MockUsersDescriberMockRecorder) UserByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByName", reflect.TypeOf((*MockUsersDescriber)(nil).UserByName), arg0)
}