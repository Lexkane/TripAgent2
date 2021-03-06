// Code generated by MockGen. DO NOT EDIT.
// Source: team-project/services/model (interfaces: Model)

// Package model is a generated GoMock package.
package model

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockModel is a mock of Model interface
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// CheckPasswordHash mocks base method
func (m *MockModel) CheckPasswordHash(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordHash", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckPasswordHash indicates an expected call of CheckPasswordHash
func (mr *MockModelMockRecorder) CheckPasswordHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordHash", reflect.TypeOf((*MockModel)(nil).CheckPasswordHash), arg0, arg1)
}

// HashPassword mocks base method
func (m *MockModel) HashPassword(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword
func (mr *MockModelMockRecorder) HashPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockModel)(nil).HashPassword), arg0)
}
