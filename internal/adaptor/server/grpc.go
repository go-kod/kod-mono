package server

import (
	"context"

	"github.com/go-kod/kod"
	snowflakev1 "github.com/go-kod/kod-mono/api/grpc/gen/go/snowflake/v1"
	"github.com/go-kod/kod-mono/internal/domain/snowflake"
	"google.golang.org/grpc/status"
)

type grpcImpl struct {
	kod.Implements[GrpcController]

	snowflake kod.Ref[snowflake.Service]
}

func (s *grpcImpl) UniqueId(ctx context.Context, req *snowflakev1.UniqueIdRequest) (*snowflakev1.UniqueIdResponse, error) {
	res, err := s.snowflake.Get().Gen(ctx, nil)
	if err != nil {
		return nil, status.Convert(err).Err()
	}

	return &snowflakev1.UniqueIdResponse{
		Uuid: res.UUID,
	}, nil
}
