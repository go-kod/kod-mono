package registry

import (
	"context"

	"google.golang.org/grpc/resolver"
)

type RegistryComponent interface {
	Register(ctx context.Context, scheme, addr string) error
	UnRegister(ctx context.Context, scheme, addr string) error
	NewBuilder(ctx context.Context) (resolver.Builder, error)
}
