// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package snowflake

import (
	"context"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod/interceptor"
	"reflect"
)

// Full method names for components.
const (
	// Service_Gen_FullMethodName is the full name of the method [service.Gen].
	Service_Gen_FullMethodName = "github.com/go-kod/kod-mono/internal/domain/snowflake/Service.Gen"
)

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod-mono/internal/domain/snowflake/Service",
		Interface: reflect.TypeOf((*Service)(nil)).Elem(),
		Impl:      reflect.TypeOf(service{}),
		Refs:      `⟦2f0f2230:KoDeDgE:github.com/go-kod/kod-mono/internal/domain/snowflake/Service→github.com/go-kod/kod-mono/internal/infra/redis/SnowflakeRepository⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return service_local_stub{
				impl:        info.Impl.(Service),
				interceptor: info.Interceptor,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[Service] = (*service)(nil)

// Local stub implementations.
// service_local_stub is a local stub implementation of [Service].
type service_local_stub struct {
	impl        Service
	interceptor interceptor.Interceptor
}

// Check that [service_local_stub] implements the [Service] interface.
var _ Service = (*service_local_stub)(nil)

// Gen wraps the method [service.Gen].
func (s service_local_stub) Gen(ctx context.Context, a1 *GenReq) (r0 *GenRes, err error) {

	if s.interceptor == nil {
		r0, err = s.impl.Gen(ctx, a1)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		r0, err = s.impl.Gen(ctx, a1)
		res[0] = r0
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: Service_Gen_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}
