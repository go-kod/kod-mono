// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/api/graph"
	"github.com/go-kod/kod-mono/api/grpc/gen/go/snowflake/v1"
	"github.com/go-kod/kod/interceptor"
	"reflect"
)

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/adaptor/server/GinController",
		Interface: reflect.TypeOf((*GinController)(nil)).Elem(),
		Impl:      reflect.TypeOf(ginImpl{}),
		Refs:      `⟦039ecf94:KoDeDgE:github.com/go-kod/kod-mono/internal/adaptor/server/GinController→github.com/go-kod/kod-mono/internal/app/example/Service⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return ginController_local_stub{
				impl:        info.Impl.(GinController),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/adaptor/server/GraphController",
		Interface: reflect.TypeOf((*GraphController)(nil)).Elem(),
		Impl:      reflect.TypeOf(resolver{}),
		Refs:      `⟦35bf9cd5:KoDeDgE:github.com/go-kod/kod-mono/internal/adaptor/server/GraphController→github.com/go-kod/kod-mono/internal/app/example/GraphService⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return graphController_local_stub{
				impl:        info.Impl.(GraphController),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/adaptor/server/GrpcController",
		Interface: reflect.TypeOf((*GrpcController)(nil)).Elem(),
		Impl:      reflect.TypeOf(grpcImpl{}),
		Refs:      `⟦74eadafa:KoDeDgE:github.com/go-kod/kod-mono/internal/adaptor/server/GrpcController→github.com/go-kod/kod-mono/internal/domain/snowflake/Service⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return grpcController_local_stub{
				impl:        info.Impl.(GrpcController),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod/Main",
		Interface: reflect.TypeOf((*kod.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(Server{}),
		Refs: `⟦3b277b6c:KoDeDgE:github.com/go-kod/kod/Main→github.com/go-kod/kod-mono/internal/adaptor/server/GinController⟧,
⟦8b746b0c:KoDeDgE:github.com/go-kod/kod/Main→github.com/go-kod/kod-mono/internal/adaptor/server/GrpcController⟧,
⟦88b98bfb:KoDeDgE:github.com/go-kod/kod/Main→github.com/go-kod/kod-mono/internal/adaptor/server/GraphController⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return main_local_stub{
				impl:        info.Impl.(kod.Main),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[GinController] = (*ginImpl)(nil)
var _ kod.InstanceOf[GraphController] = (*resolver)(nil)
var _ kod.InstanceOf[GrpcController] = (*grpcImpl)(nil)
var _ kod.InstanceOf[kod.Main] = (*Server)(nil)

// Local stub implementations.

// ginController_local_stub is a local stub implementation of [GinController].
type ginController_local_stub struct {
	impl        GinController
	name        string
	interceptor interceptor.Interceptor
}

// Check that ginController_local_stub implements the GinController interface.
var _ GinController = (*ginController_local_stub)(nil)

func (s ginController_local_stub) UniqueID(a0 *gin.Context) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	s.impl.UniqueID(a0)
	return
}

// graphController_local_stub is a local stub implementation of [GraphController].
type graphController_local_stub struct {
	impl        GraphController
	name        string
	interceptor interceptor.Interceptor
}

// Check that graphController_local_stub implements the GraphController interface.
var _ GraphController = (*graphController_local_stub)(nil)

func (s graphController_local_stub) Mutation() (r0 graph.MutationResolver) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.Mutation()
	return
}

func (s graphController_local_stub) Query() (r0 graph.QueryResolver) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.Query()
	return
}

// grpcController_local_stub is a local stub implementation of [GrpcController].
type grpcController_local_stub struct {
	impl        GrpcController
	name        string
	interceptor interceptor.Interceptor
}

// Check that grpcController_local_stub implements the GrpcController interface.
var _ GrpcController = (*grpcController_local_stub)(nil)

func (s grpcController_local_stub) UniqueId(ctx context.Context, a1 *snowflakev1.UniqueIdRequest) (r0 *snowflakev1.UniqueIdResponse, err error) {

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
		FullMethod: "github.com/go-kod/kod-mono/internal/adaptor/server/GrpcController.UniqueId",
		Method:     "UniqueId",
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}

// main_local_stub is a local stub implementation of [kod.Main].
type main_local_stub struct {
	impl        kod.Main
	name        string
	interceptor interceptor.Interceptor
}

// Check that main_local_stub implements the kod.Main interface.
var _ kod.Main = (*main_local_stub)(nil)

