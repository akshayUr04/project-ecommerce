// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/akshayur04/project-ecommerce/pkg/repository/interface (interfaces: UserRepository)

// Package mockRepo is a generated GoMock package.
package mockRepo

import (
	context "context"
	reflect "reflect"

	helperStruct "github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	response "github.com/akshayur04/project-ecommerce/pkg/common/response"
	domain "github.com/akshayur04/project-ecommerce/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(arg0 int, arg1 helperStruct.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), arg0, arg1)
}

// FindPassword mocks base method.
func (m *MockUserRepository) FindPassword(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPassword", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPassword indicates an expected call of FindPassword.
func (mr *MockUserRepositoryMockRecorder) FindPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPassword", reflect.TypeOf((*MockUserRepository)(nil).FindPassword), arg0)
}

// IsSignIn mocks base method.
func (m *MockUserRepository) IsSignIn(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSignIn", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSignIn indicates an expected call of IsSignIn.
func (mr *MockUserRepositoryMockRecorder) IsSignIn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSignIn", reflect.TypeOf((*MockUserRepository)(nil).IsSignIn), arg0)
}

// OtpLogin mocks base method.
func (m *MockUserRepository) OtpLogin(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OtpLogin", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OtpLogin indicates an expected call of OtpLogin.
func (mr *MockUserRepositoryMockRecorder) OtpLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OtpLogin", reflect.TypeOf((*MockUserRepository)(nil).OtpLogin), arg0)
}

// UpdateAddress mocks base method.
func (m *MockUserRepository) UpdateAddress(arg0, arg1 int, arg2 helperStruct.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserRepositoryMockRecorder) UpdateAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserRepository)(nil).UpdateAddress), arg0, arg1, arg2)
}

// UpdatePassword mocks base method.
func (m *MockUserRepository) UpdatePassword(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserRepositoryMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserRepository)(nil).UpdatePassword), arg0, arg1)
}

// UserEditProfile mocks base method.
func (m *MockUserRepository) UserEditProfile(arg0 int, arg1 helperStruct.UserReq) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserEditProfile", arg0, arg1)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserEditProfile indicates an expected call of UserEditProfile.
func (mr *MockUserRepositoryMockRecorder) UserEditProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserEditProfile", reflect.TypeOf((*MockUserRepository)(nil).UserEditProfile), arg0, arg1)
}

// UserLogin mocks base method.
func (m *MockUserRepository) UserLogin(arg0 context.Context, arg1 string) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", arg0, arg1)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockUserRepositoryMockRecorder) UserLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockUserRepository)(nil).UserLogin), arg0, arg1)
}

// UserSignUp mocks base method.
func (m *MockUserRepository) UserSignUp(arg0 context.Context, arg1 helperStruct.UserReq) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", arg0, arg1)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserRepositoryMockRecorder) UserSignUp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserRepository)(nil).UserSignUp), arg0, arg1)
}

// Viewprfile mocks base method.
func (m *MockUserRepository) Viewprfile(arg0 int) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Viewprfile", arg0)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Viewprfile indicates an expected call of Viewprfile.
func (mr *MockUserRepositoryMockRecorder) Viewprfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Viewprfile", reflect.TypeOf((*MockUserRepository)(nil).Viewprfile), arg0)
}