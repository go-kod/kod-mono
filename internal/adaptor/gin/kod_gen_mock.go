// Code generated by MockGen. DO NOT EDIT.
// Source: internal/adaptor/gin/kod_gen_interface.go
//
// Generated by this command:
//
//	mockgen -source internal/adaptor/gin/kod_gen_interface.go -destination internal/adaptor/gin/kod_gen_mock.go -package gin
//
// Package gin is a generated GoMock package.
package gin

import (
	reflect "reflect"

	kgin "github.com/go-kod/kod-ext/server/kgin"
	gomock "go.uber.org/mock/gomock"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// UniqueID mocks base method.
func (m *MockController) UniqueID(ctx *kgin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UniqueID", ctx)
}

// UniqueID indicates an expected call of UniqueID.
func (mr *MockControllerMockRecorder) UniqueID(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniqueID", reflect.TypeOf((*MockController)(nil).UniqueID), ctx)
}
