package example

import (
	"context"
	"log/slog"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/internal/domain/snowflake"
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

	uuidService kod.Ref[snowflake.Service]
}

func (c *component) UniqueID(ctx context.Context, req *TestReq) (*TestRes, error) {
	c.L().InfoContext(ctx, "UniqueID", slog.Any("req", req))
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is empty")
	}

	res, err := c.uuidService.Get().Gen(ctx, &snowflake.GenReq{})
	if err != nil {
		return nil, err
	}

	return &TestRes{Uuid: res.UUID}, nil
}
