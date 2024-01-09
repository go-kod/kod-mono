// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/example/kod_gen_interface.go
//
// Generated by this command:
//
//	mockgen -source internal/app/example/kod_gen_interface.go -destination internal/app/example/kod_gen_mock.go -package example
//
// Package example is a generated GoMock package.
package example

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
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

// UniqueID mocks base method.
func (m *MockService) UniqueID(ctx context.Context, req *TestReq) (*TestRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UniqueID", ctx, req)
	ret0, _ := ret[0].(*TestRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UniqueID indicates an expected call of UniqueID.
func (mr *MockServiceMockRecorder) UniqueID(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniqueID", reflect.TypeOf((*MockService)(nil).UniqueID), ctx, req)
}
