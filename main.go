package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-ext/client/kpyroscope"
	"github.com/go-kod/kod-ext/client/kuptrace"
	"github.com/go-kod/kod-ext/registry/etcdv3"
	kgin "github.com/go-kod/kod-ext/server/kgin"
	"github.com/go-kod/kod-mono/internal/adaptor/gin"
	"github.com/grafana/pyroscope-go"
	"github.com/samber/lo"
)

type app struct {
	kod.Implements[kod.Main]
	kod.WithConfig[config]

	server    *kgin.Server
	pyroscope *pyroscope.Profiler
	uptrace   *kuptrace.Client

	gin kod.Ref[gin.Controller]
}

func (s *app) Init(ctx context.Context) error {
	s.uptrace = lo.Must(s.Config().Uptrace.Build(ctx))

	s.pyroscope = lo.Must(s.Config().Pyroscope.Build(ctx))

	s.server = s.Config().HTTP.Build().WithRegistry(lo.Must(s.Config().Etcdv3.Build(ctx)))
	gin.Register(s.server, s.gin.Get())

	return nil
}

func (s *app) Stop(ctx context.Context) error {
	err := s.uptrace.Stop(ctx)
	if err != nil {
		return fmt.Errorf("failed to stop uptrace: %w", err)
	}

	err = s.pyroscope.Stop()
	if err != nil {
		return fmt.Errorf("failed to stop pyroscope: %w", err)
	}

	err = s.server.GracefulStop(ctx)
	if err != nil {
		return fmt.Errorf("failed to stop server: %w", err)
	}

	return nil
}

type config struct {
	HTTP      kgin.Config       `toml:"http"`
	Uptrace   kuptrace.Config   `toml:"uptrace"`
	Pyroscope kpyroscope.Config `toml:"pyroscope"`
	Etcdv3    etcdv3.Config     `toml:"etcdv3"`
}

//	@title			Swagger Example API
//	@version		2.0
//	@description	This is a sample server.
//	@termsOfService	http://swagger.io/terms/
//	@host			localhost:9257
//	@schemes		http

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	/
func main() {
	lo.Must0(kod.Run(context.Background(), func(ctx context.Context, s *app) error {
		err := s.server.Run(ctx)
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return errors.Join(err)
	}))
}
