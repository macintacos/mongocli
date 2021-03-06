// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: EventLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockEventLister is a mock of EventLister interface
type MockEventLister struct {
	ctrl     *gomock.Controller
	recorder *MockEventListerMockRecorder
}

// MockEventListerMockRecorder is the mock recorder for MockEventLister
type MockEventListerMockRecorder struct {
	mock *MockEventLister
}

// NewMockEventLister creates a new mock instance
func NewMockEventLister(ctrl *gomock.Controller) *MockEventLister {
	mock := &MockEventLister{ctrl: ctrl}
	mock.recorder = &MockEventListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventLister) EXPECT() *MockEventListerMockRecorder {
	return m.recorder
}

// OrganizationEvents mocks base method
func (m *MockEventLister) OrganizationEvents(arg0 string, arg1 *mongodbatlas.EventListOptions) (*mongodbatlas.EventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationEvents", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationEvents indicates an expected call of OrganizationEvents
func (mr *MockEventListerMockRecorder) OrganizationEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationEvents", reflect.TypeOf((*MockEventLister)(nil).OrganizationEvents), arg0, arg1)
}

// ProjectEvents mocks base method
func (m *MockEventLister) ProjectEvents(arg0 string, arg1 *mongodbatlas.EventListOptions) (*mongodbatlas.EventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectEvents", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectEvents indicates an expected call of ProjectEvents
func (mr *MockEventListerMockRecorder) ProjectEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectEvents", reflect.TypeOf((*MockEventLister)(nil).ProjectEvents), arg0, arg1)
}
