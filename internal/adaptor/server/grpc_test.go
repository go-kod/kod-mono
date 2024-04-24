package server

import (
	"context"
	"testing"

	"github.com/go-kod/kod"
	snowflakev1 "github.com/go-kod/kod-mono/api/gen/go/snowflake/v1"
	"github.com/stretchr/testify/require"
)

func TestGrpcController(t *testing.T) {
	kod.RunTest(t, func(ctx context.Context, c GrpcController) {
		res, err := c.UniqueId(ctx, &snowflakev1.UniqueIdRequest{})
		require.Nil(t, err)
		require.NotEmpty(t, res.GetUuid())
	}, kod.WithConfigFile("../../../config/server/dev.toml"))
}