package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod/ext/client/kpyroscope"
	"github.com/go-kod/kod/ext/client/kuptrace"
	"github.com/go-kod/kod/ext/registry/etcdv3"
	kgin "github.com/go-kod/kod/ext/server/kgin"
	"github.com/grafana/pyroscope-go"
	"github.com/samber/lo"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	kod.Implements[kod.Main]
	kod.WithConfig[config]

	server    *kgin.Server
	pyroscope *pyroscope.Profiler
	uptrace   *kuptrace.Client

	example kod.Ref[Controller]
}

func (s *Server) Init(ctx context.Context) error {
	s.uptrace = lo.Must(s.Config().Uptrace.Build(ctx))

	s.pyroscope = lo.Must(s.Config().Pyroscope.Build(ctx))

	s.server = s.Config().HTTP.Build().WithRegistry(lo.Must(s.Config().Etcdv3.Build(ctx)))
	Register(s.server, s.example.Get())

	// Swagger
	s.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	return nil
}

func (s *Server) Run(ctx context.Context) error {
	err := s.server.Run(ctx)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
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
