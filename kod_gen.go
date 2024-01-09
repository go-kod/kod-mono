// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package main

import (
	"context"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod/core/interceptor"
	"reflect"
)

func init() {
	kod.Register(kod.Registration{
		Name:  "github.com/go-kod/kod/Main",
		Iface: reflect.TypeOf((*kod.Main)(nil)).Elem(),
		Impl:  reflect.TypeOf(app{}),
		Refs:  `⟦7446e71c:KoDeDgE:github.com/go-kod/kod/Main→github.com/go-kod/kod-mono/internal/adaptor/gin/Controller⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			var interceptors []kod.Interceptor
			if h, ok := info.Impl.(interface{ Interceptors() []kod.Interceptor }); ok {
				interceptors = h.Interceptors()
			}

			return main_local_stub{
				impl:        info.Impl.(kod.Main),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
				caller:      info.Caller,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[kod.Main] = (*app)(nil)

// Local stub implementations.

type main_local_stub struct {
	impl        kod.Main
	name        string
	caller      string
	interceptor kod.Interceptor
}

// Check that main_local_stub implements the kod.Main interface.
var _ kod.Main = (*main_local_stub)(nil)

