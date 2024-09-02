// Code generated by MockGen. DO NOT EDIT.
// Source: task.go

// Package mocktask is a generated GoMock package.
package mocktask

import (
	entity "dddapib/internal/domain/model/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CloseTask mocks base method.
func (m *MockService) CloseTask(id string) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseTask", id)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseTask indicates an expected call of CloseTask.
func (mr *MockServiceMockRecorder) CloseTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseTask", reflect.TypeOf((*MockService)(nil).CloseTask), id)
}

// CreateTask mocks base method.
func (m *MockService) CreateTask(name string) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", name)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockServiceMockRecorder) CreateTask(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockService)(nil).CreateTask), name)
}

// DeleteTask mocks base method.
func (m *MockService) DeleteTask(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockServiceMockRecorder) DeleteTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockService)(nil).DeleteTask), id)
}

// ListTasks mocks base method.
func (m *MockService) ListTasks() ([]*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks")
	ret0, _ := ret[0].([]*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks.
func (mr *MockServiceMockRecorder) ListTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockService)(nil).ListTasks))
}
