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

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/app/example/GraphService",
		Interface: reflect.TypeOf((*GraphService)(nil)).Elem(),
		Impl:      reflect.TypeOf(graphImpl{}),
		Refs:      ``,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return graphService_local_stub{
				impl:        info.Impl.(GraphService),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/app/example/Service",
		Interface: reflect.TypeOf((*Service)(nil)).Elem(),
		Impl:      reflect.TypeOf(component{}),
		Refs:      `⟦2cd73ff4:KoDeDgE:github.com/go-kod/kod-mono/internal/app/example/Service→github.com/go-kod/kod-mono/internal/domain/snowflake/Service⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return service_local_stub{
				impl:        info.Impl.(Service),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[GraphService] = (*graphImpl)(nil)
var _ kod.InstanceOf[Service] = (*component)(nil)

// Local stub implementations.

type graphService_local_stub struct {
	impl        GraphService
	name        string
	interceptor interceptor.Interceptor
}

// Check that graphService_local_stub implements the GraphService interface.
var _ GraphService = (*graphService_local_stub)(nil)

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
		Component:  s.name,
		FullMethod: "github.com/go-kod/kod-mono/internal/app/example/GraphService.CreateTodo",
		Method:     "CreateTodo",
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}

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
		Component:  s.name,
		FullMethod: "github.com/go-kod/kod-mono/internal/app/example/GraphService.Todos",
		Method:     "Todos",
	}

	err = s.interceptor(ctx, info, []any{}, []any{r0}, call)
	return
}

type service_local_stub struct {
	impl        Service
	name        string
	interceptor interceptor.Interceptor
}

// Check that service_local_stub implements the Service interface.
var _ Service = (*service_local_stub)(nil)

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
		Component:  s.name,
		FullMethod: "github.com/go-kod/kod-mono/internal/app/example/Service.UniqueID",
		Method:     "UniqueID",
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}
