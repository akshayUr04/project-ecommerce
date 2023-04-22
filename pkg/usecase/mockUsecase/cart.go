// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/akshayur04/project-ecommerce/pkg/usecase/interface (interfaces: CartUsecase)

// Package mockUsecase is a generated GoMock package.
package mockUsecase

import (
	reflect "reflect"

	response "github.com/akshayur04/project-ecommerce/pkg/common/response"
	gomock "github.com/golang/mock/gomock"
)

// MockCartUsecase is a mock of CartUsecase interface.
type MockCartUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCartUsecaseMockRecorder
}

// MockCartUsecaseMockRecorder is the mock recorder for MockCartUsecase.
type MockCartUsecaseMockRecorder struct {
	mock *MockCartUsecase
}

// NewMockCartUsecase creates a new mock instance.
func NewMockCartUsecase(ctrl *gomock.Controller) *MockCartUsecase {
	mock := &MockCartUsecase{ctrl: ctrl}
	mock.recorder = &MockCartUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartUsecase) EXPECT() *MockCartUsecaseMockRecorder {
	return m.recorder
}

// AddToCart mocks base method.
func (m *MockCartUsecase) AddToCart(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToCart indicates an expected call of AddToCart.
func (mr *MockCartUsecaseMockRecorder) AddToCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCart", reflect.TypeOf((*MockCartUsecase)(nil).AddToCart), arg0, arg1)
}

// CreateCart mocks base method.
func (m *MockCartUsecase) CreateCart(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockCartUsecaseMockRecorder) CreateCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockCartUsecase)(nil).CreateCart), arg0)
}

// ListCart mocks base method.
func (m *MockCartUsecase) ListCart(arg0 int) (response.ViewCart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCart", arg0)
	ret0, _ := ret[0].(response.ViewCart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCart indicates an expected call of ListCart.
func (mr *MockCartUsecaseMockRecorder) ListCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCart", reflect.TypeOf((*MockCartUsecase)(nil).ListCart), arg0)
}

// RemoveFromCart mocks base method.
func (m *MockCartUsecase) RemoveFromCart(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromCart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromCart indicates an expected call of RemoveFromCart.
func (mr *MockCartUsecaseMockRecorder) RemoveFromCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromCart", reflect.TypeOf((*MockCartUsecase)(nil).RemoveFromCart), arg0, arg1)
}
