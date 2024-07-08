package grpc

import (
	"context"

	"github.com/go-kod/kod"
	snowflakev1 "github.com/go-kod/kod-mono/api/grpc/gen/go/snowflake/v1"
	"github.com/go-kod/kod/ext/client/kgrpc"
)

type impl struct {
	kod.Implements[Snowflake]
	kod.WithConfig[kgrpc.Config]

	cc snowflakev1.SnowflakeServiceClient
}

func (impl *impl) Init(ctx context.Context) error {
	impl.cc = snowflakev1.NewSnowflakeServiceClient(impl.Config().Build())

	return nil
}

func (impl *impl) UniqueId(ctx context.Context, req *snowflakev1.UniqueIdRequest) (*snowflakev1.UniqueIdResponse, error) {
	return impl.cc.UniqueId(ctx, req)
}
