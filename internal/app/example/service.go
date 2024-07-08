package example

import (
	"context"
	"log/slog"

	"github.com/go-kod/kod"
	snowflakev1 "github.com/go-kod/kod-mono/api/grpc/gen/go/snowflake/v1"
	"github.com/go-kod/kod-mono/internal/infra/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TestReq ...
type TestReq struct {
	Name string `form:"name"`
}

// TestRes ...
type TestRes struct {
	// 唯一标识
	Uuid string
}

type component struct {
	kod.Implements[Service]

	grpcSnowflake kod.Ref[grpc.Snowflake]
}

func (c *component) UniqueID(ctx context.Context, req *TestReq) (*TestRes, error) {
	c.L(ctx).InfoContext(ctx, "UniqueID", slog.Any("req", req))
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is empty")
	}

	res, err := c.grpcSnowflake.Get().UniqueId(ctx, &snowflakev1.UniqueIdRequest{})
	if err != nil {
		return nil, err
	}

	return &TestRes{Uuid: res.GetUuid()}, nil
}
