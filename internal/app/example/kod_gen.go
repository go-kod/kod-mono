// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package example

import (
	"context"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/api/graph/model"
	"github.com/go-kod/kod/interceptor"
	"reflect"
)

// Full method names for components.
const (
	// GraphService_CreateTodo_FullMethodName is the full name of the method [graphImpl.CreateTodo].
	GraphService_CreateTodo_FullMethodName = "github.com/go-kod/kod-mono/internal/app/example/GraphService.CreateTodo"
	// GraphService_Todos_FullMethodName is the full name of the method [graphImpl.Todos].
	GraphService_Todos_FullMethodName = "github.com/go-kod/kod-mono/internal/app/example/GraphService.Todos"
	// Service_UniqueID_FullMethodName is the full name of the method [component.UniqueID].
	Service_UniqueID_FullMethodName = "github.com/go-kod/kod-mono/internal/app/example/Service.UniqueID"
)

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/app/example/GraphService",
		Interface: reflect.TypeOf((*GraphService)(nil)).Elem(),
		Impl:      reflect.TypeOf(graphImpl{}),
		Refs:      ``,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return graphService_local_stub{
				impl:        info.Impl.(GraphService),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/app/example/Service",
		Interface: reflect.TypeOf((*Service)(nil)).Elem(),
		Impl:      reflect.TypeOf(component{}),
		Refs:      `⟦2c62dafb:KoDeDgE:github.com/go-kod/kod-mono/internal/app/example/Service→github.com/go-kod/kod-mono/internal/infra/grpc/Snowflake⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return service_local_stub{
				impl:        info.Impl.(Service),
				interceptor: info.Interceptor,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[GraphService] = (*graphImpl)(nil)
var _ kod.InstanceOf[Service] = (*component)(nil)

// Local stub implementations.

// graphService_local_stub is a local stub implementation of [GraphService].
type graphService_local_stub struct {
	impl        GraphService
	interceptor interceptor.Interceptor
}

// Check that [graphService_local_stub] implements the [GraphService] interface.
var _ GraphService = (*graphService_local_stub)(nil)

// CreateTodo wraps the method [graphImpl.CreateTodo].
func (s graphService_local_stub) CreateTodo(ctx context.Context, a1 model.NewTodo) (r0 *model.Todo, err error) {

	if s.interceptor == nil {
		r0, err = s.impl.CreateTodo(ctx, a1)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		r0, err = s.impl.CreateTodo(ctx, a1)
		res[0] = r0
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: GraphService_CreateTodo_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}

// Todos wraps the method [graphImpl.Todos].
func (s graphService_local_stub) Todos(ctx context.Context) (r0 []*model.Todo, err error) {

	if s.interceptor == nil {
		r0, err = s.impl.Todos(ctx)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		r0, err = s.impl.Todos(ctx)
		res[0] = r0
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: GraphService_Todos_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{}, []any{r0}, call)
	return
}

// service_local_stub is a local stub implementation of [Service].
type service_local_stub struct {
	impl        Service
	interceptor interceptor.Interceptor
}

// Check that [service_local_stub] implements the [Service] interface.
var _ Service = (*service_local_stub)(nil)

// UniqueID wraps the method [component.UniqueID].
func (s service_local_stub) UniqueID(ctx context.Context, a1 *TestReq) (r0 *TestRes, err error) {

	if s.interceptor == nil {
		r0, err = s.impl.UniqueID(ctx, a1)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		r0, err = s.impl.UniqueID(ctx, a1)
		res[0] = r0
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: Service_UniqueID_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}
