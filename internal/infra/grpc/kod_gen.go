// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package grpc

import (
	"context"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/api/grpc/gen/go/snowflake/v1"
	"github.com/go-kod/kod/interceptor"
	"reflect"
)

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/infra/grpc/Snowflake",
		Interface: reflect.TypeOf((*Snowflake)(nil)).Elem(),
		Impl:      reflect.TypeOf(impl{}),
		Refs:      ``,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return snowflake_local_stub{
				impl:        info.Impl.(Snowflake),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[Snowflake] = (*impl)(nil)

// Local stub implementations.

// snowflake_local_stub is a local stub implementation of [Snowflake].
type snowflake_local_stub struct {
	impl        Snowflake
	name        string
	interceptor interceptor.Interceptor
}

// Check that snowflake_local_stub implements the Snowflake interface.
var _ Snowflake = (*snowflake_local_stub)(nil)

func (s snowflake_local_stub) UniqueId(ctx context.Context, a1 *snowflakev1.UniqueIdRequest) (r0 *snowflakev1.UniqueIdResponse, err error) {

	if s.interceptor == nil {
		r0, err = s.impl.UniqueId(ctx, a1)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		r0, err = s.impl.UniqueId(ctx, a1)
		res[0] = r0
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		Component:  s.name,
		FullMethod: "github.com/go-kod/kod-mono/internal/infra/grpc/Snowflake.UniqueId",
		Method:     "UniqueId",
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}
