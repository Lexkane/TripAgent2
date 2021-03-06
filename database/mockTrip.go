// Code generated by MockGen. DO NOT EDIT.
// Source: team-project/database (interfaces: TripRepository)

// Package database is a generated GoMock package.
package database

import (
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
	data "team-project/services/data"
)

// MockTripRepository is a mock of TripRepository interface
type MockTripRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTripRepositoryMockRecorder
}

// MockTripRepositoryMockRecorder is the mock recorder for MockTripRepository
type MockTripRepositoryMockRecorder struct {
	mock *MockTripRepository
}

// NewMockTripRepository creates a new mock instance
func NewMockTripRepository(ctrl *gomock.Controller) *MockTripRepository {
	mock := &MockTripRepository{ctrl: ctrl}
	mock.recorder = &MockTripRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTripRepository) EXPECT() *MockTripRepositoryMockRecorder {
	return m.recorder
}

// AddTrip mocks base method
func (m *MockTripRepository) AddTrip(arg0 data.Trip) (data.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrip", arg0)
	ret0, _ := ret[0].(data.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTrip indicates an expected call of AddTrip
func (mr *MockTripRepositoryMockRecorder) AddTrip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrip", reflect.TypeOf((*MockTripRepository)(nil).AddTrip), arg0)
}

// DeleteTrip mocks base method
func (m *MockTripRepository) DeleteTrip(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrip", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrip indicates an expected call of DeleteTrip
func (mr *MockTripRepositoryMockRecorder) DeleteTrip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrip", reflect.TypeOf((*MockTripRepository)(nil).DeleteTrip), arg0)
}

// GetTrip mocks base method
func (m *MockTripRepository) GetTrip(arg0 uuid.UUID) (data.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrip", arg0)
	ret0, _ := ret[0].(data.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrip indicates an expected call of GetTrip
func (mr *MockTripRepositoryMockRecorder) GetTrip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrip", reflect.TypeOf((*MockTripRepository)(nil).GetTrip), arg0)
}

// GetTrips mocks base method
func (m *MockTripRepository) GetTrips() ([]data.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrips")
	ret0, _ := ret[0].([]data.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrips indicates an expected call of GetTrips
func (mr *MockTripRepositoryMockRecorder) GetTrips() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrips", reflect.TypeOf((*MockTripRepository)(nil).GetTrips))
}

// UpdateTrip mocks base method
func (m *MockTripRepository) UpdateTrip(arg0 data.Trip) (data.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrip", arg0)
	ret0, _ := ret[0].(data.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrip indicates an expected call of UpdateTrip
func (mr *MockTripRepositoryMockRecorder) UpdateTrip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrip", reflect.TypeOf((*MockTripRepository)(nil).UpdateTrip), arg0)
}
