package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/internal/adaptor/server"
	"github.com/stretchr/testify/require"
)

func TestGinController(t *testing.T) {
	kod.RunTest(t, func(ctx context.Context, s server.GinController) {
		record := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(record)
		c.Request, _ = http.NewRequest("GET", "/uniqueId?name=bob", nil)

		s.UniqueID(c)

		require.Equal(t, 200, record.Code)
		require.Equal(t, "application/json; charset=utf-8", record.Header().Get("Content-Type"))
		require.NotEmpty(t, record.Body.String())
	}, kod.WithConfigFile("../../../config/server/dev.toml"))
}
