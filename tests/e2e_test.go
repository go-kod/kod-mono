package tests

import (
	"context"
	"testing"
	"time"

	_ "github.com/go-kod/kod-mono/tests/server"
	"github.com/stretchr/testify/assert"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/internal/adaptor/server"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"
)

func TestE2ESuites(t *testing.T) {
	err := kod.Run(context.Background(), func(ctx context.Context, s *server.Server) error {
		go func() {
			lo.Must0(s.Run(ctx))
		}()

		// wait for server start
		time.Sleep(time.Second)
		RegisterFailHandler(Fail)
		RunSpecs(t, "mock test cases")

		return nil
	}, kod.WithConfigFile("../config/server/dev.toml"))

	assert.Nil(t, err)
}
