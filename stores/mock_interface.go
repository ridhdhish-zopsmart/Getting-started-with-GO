// Code generated by MockGen. DO NOT EDIT.
// Source: layer/user/stores (interfaces: User)

// Package stores is a generated GoMock package.
package stores

import (
	gomock "github.com/golang/mock/gomock"
	models "layer/user/models"
	reflect "reflect"
)

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// GetUserById mocks base method
func (m *MockUser) GetUserById(arg0 int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockUserMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUser)(nil).GetUserById), arg0)
}

// GetUsers mocks base method
func (m *MockUser) GetUsers() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockUserMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUser)(nil).GetUsers))
}

// UpdateUser mocks base method
func (m *MockUser) UpdateUser(arg0 int, arg1 models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), arg0, arg1)
}
