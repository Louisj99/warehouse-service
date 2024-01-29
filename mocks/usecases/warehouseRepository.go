// Code generated by MockGen. DO NOT EDIT.
// Source: warehouse-service/pkg/usecases (interfaces: WarehouseRepository)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination=../../mocks/usecases/warehouseRepository.go . WarehouseRepository
//
// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"
	entities "warehouse-service/pkg/entities"

	gomock "go.uber.org/mock/gomock"
)

// MockWarehouseRepository is a mock of WarehouseRepository interface.
type MockWarehouseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWarehouseRepositoryMockRecorder
}

// MockWarehouseRepositoryMockRecorder is the mock recorder for MockWarehouseRepository.
type MockWarehouseRepositoryMockRecorder struct {
	mock *MockWarehouseRepository
}

// NewMockWarehouseRepository creates a new mock instance.
func NewMockWarehouseRepository(ctrl *gomock.Controller) *MockWarehouseRepository {
	mock := &MockWarehouseRepository{ctrl: ctrl}
	mock.recorder = &MockWarehouseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWarehouseRepository) EXPECT() *MockWarehouseRepositoryMockRecorder {
	return m.recorder
}

// GetAllItems mocks base method.
func (m *MockWarehouseRepository) GetAllItems(arg0 context.Context) ([]entities.ItemInformation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllItems", arg0)
	ret0, _ := ret[0].([]entities.ItemInformation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllItems indicates an expected call of GetAllItems.
func (mr *MockWarehouseRepositoryMockRecorder) GetAllItems(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllItems", reflect.TypeOf((*MockWarehouseRepository)(nil).GetAllItems), arg0)
}

// GetItemInformation mocks base method.
func (m *MockWarehouseRepository) GetItemInformation(arg0 context.Context, arg1 string) (*entities.ItemInformation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemInformation", arg0, arg1)
	ret0, _ := ret[0].(*entities.ItemInformation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemInformation indicates an expected call of GetItemInformation.
func (mr *MockWarehouseRepositoryMockRecorder) GetItemInformation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemInformation", reflect.TypeOf((*MockWarehouseRepository)(nil).GetItemInformation), arg0, arg1)
}

// PlaceholderAdapter mocks base method.
func (m *MockWarehouseRepository) PlaceholderAdapter(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlaceholderAdapter", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PlaceholderAdapter indicates an expected call of PlaceholderAdapter.
func (mr *MockWarehouseRepositoryMockRecorder) PlaceholderAdapter(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlaceholderAdapter", reflect.TypeOf((*MockWarehouseRepository)(nil).PlaceholderAdapter), arg0, arg1)
}

// UpdateItemInformation mocks base method.
func (m *MockWarehouseRepository) UpdateItemInformation(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemInformation", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItemInformation indicates an expected call of UpdateItemInformation.
func (mr *MockWarehouseRepositoryMockRecorder) UpdateItemInformation(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemInformation", reflect.TypeOf((*MockWarehouseRepository)(nil).UpdateItemInformation), arg0, arg1, arg2, arg3, arg4, arg5)
}
