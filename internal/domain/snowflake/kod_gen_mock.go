// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/snowflake/kod_gen_interface.go
//
// Generated by this command:
//
//	mockgen -source internal/domain/snowflake/kod_gen_interface.go -destination internal/domain/snowflake/kod_gen_mock.go -package snowflake -typed
//

// Package snowflake is a generated GoMock package.
package snowflake

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

// Gen mocks base method.
func (m *MockService) Gen(ctx context.Context, arg1 *GenReq) (*GenRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gen", ctx, arg1)
	ret0, _ := ret[0].(*GenRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Gen indicates an expected call of Gen.
func (mr *MockServiceMockRecorder) Gen(ctx, arg1 any) *MockServiceGenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gen", reflect.TypeOf((*MockService)(nil).Gen), ctx, arg1)
	return &MockServiceGenCall{Call: call}
}

// MockServiceGenCall wrap *gomock.Call
type MockServiceGenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceGenCall) Return(arg0 *GenRes, arg1 error) *MockServiceGenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceGenCall) Do(f func(context.Context, *GenReq) (*GenRes, error)) *MockServiceGenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceGenCall) DoAndReturn(f func(context.Context, *GenReq) (*GenRes, error)) *MockServiceGenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
