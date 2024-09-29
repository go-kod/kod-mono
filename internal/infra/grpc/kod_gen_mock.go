// Code generated by MockGen. DO NOT EDIT.
// Source: internal/infra/grpc/kod_gen_interface.go
//
// Generated by this command:
//
//	mockgen -source internal/infra/grpc/kod_gen_interface.go -destination internal/infra/grpc/kod_gen_mock.go -package grpc -typed
//

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	reflect "reflect"

	snowflakev1 "github.com/go-kod/kod-mono/api/grpc/gen/go/snowflake/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockSnowflake is a mock of Snowflake interface.
type MockSnowflake struct {
	ctrl     *gomock.Controller
	recorder *MockSnowflakeMockRecorder
}

// MockSnowflakeMockRecorder is the mock recorder for MockSnowflake.
type MockSnowflakeMockRecorder struct {
	mock *MockSnowflake
}

// NewMockSnowflake creates a new mock instance.
func NewMockSnowflake(ctrl *gomock.Controller) *MockSnowflake {
	mock := &MockSnowflake{ctrl: ctrl}
	mock.recorder = &MockSnowflakeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnowflake) EXPECT() *MockSnowflakeMockRecorder {
	return m.recorder
}

// UniqueId mocks base method.
func (m *MockSnowflake) UniqueId(ctx context.Context, req *snowflakev1.UniqueIdRequest) (*snowflakev1.UniqueIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UniqueId", ctx, req)
	ret0, _ := ret[0].(*snowflakev1.UniqueIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UniqueId indicates an expected call of UniqueId.
func (mr *MockSnowflakeMockRecorder) UniqueId(ctx, req any) *MockSnowflakeUniqueIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniqueId", reflect.TypeOf((*MockSnowflake)(nil).UniqueId), ctx, req)
	return &MockSnowflakeUniqueIdCall{Call: call}
}

// MockSnowflakeUniqueIdCall wrap *gomock.Call
type MockSnowflakeUniqueIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnowflakeUniqueIdCall) Return(arg0 *snowflakev1.UniqueIdResponse, arg1 error) *MockSnowflakeUniqueIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnowflakeUniqueIdCall) Do(f func(context.Context, *snowflakev1.UniqueIdRequest) (*snowflakev1.UniqueIdResponse, error)) *MockSnowflakeUniqueIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnowflakeUniqueIdCall) DoAndReturn(f func(context.Context, *snowflakev1.UniqueIdRequest) (*snowflakev1.UniqueIdResponse, error)) *MockSnowflakeUniqueIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
