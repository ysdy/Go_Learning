// Code generated by MockGen. DO NOT EDIT.
// Source: service/pets_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/ysdy/Go_Learning/entity"
	reflect "reflect"
)

// MockPetService is a mock of PetService interface.
type MockPetService struct {
	ctrl     *gomock.Controller
	recorder *MockPetServiceMockRecorder
}

// MockPetServiceMockRecorder is the mock recorder for MockPetService.
type MockPetServiceMockRecorder struct {
	mock *MockPetService
}

// NewMockPetService creates a new mock instance.
func NewMockPetService(ctrl *gomock.Controller) *MockPetService {
	mock := &MockPetService{ctrl: ctrl}
	mock.recorder = &MockPetServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPetService) EXPECT() *MockPetServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockPetService) List() ([]entity.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]entity.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPetServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPetService)(nil).List))
}
