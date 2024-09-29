package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-ext/client/kpyroscope"
	"github.com/go-kod/kod-ext/registry/etcdv3"
	kgin "github.com/go-kod/kod-ext/server/kgin"
	"github.com/go-kod/kod-ext/server/kgrpc"
	snowflakev1 "github.com/go-kod/kod-mono/api/grpc/gen/go/snowflake/v1"
	"github.com/grafana/pyroscope-go"
	"github.com/samber/lo"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	kod.Implements[kod.Main]
	kod.WithConfig[config]

	server    *kgin.Server
	grpc      *kgrpc.Server
	pyroscope *pyroscope.Profiler

	example   kod.Ref[GinController]
	grpcImpl  kod.Ref[GrpcController]
	graphImpl kod.Ref[GraphController]
}

func (s *Server) Init(ctx context.Context) error {
	s.pyroscope = lo.Must(s.Config().Pyroscope.Build(ctx))

	registry := lo.Must(s.Config().Etcdv3.Build(ctx))

	s.server = s.Config().HTTP.Build().WithRegistry(registry)
	registerHTTP(s.server, s.example.Get(), s.graphImpl.Get())

	// Swagger
	s.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	s.grpc = s.Config().Grpc.Build().WithRegistry(registry)
	snowflakev1.RegisterSnowflakeServiceServer(s.grpc, s.grpcImpl.Get())

	return nil
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		lo.Must0(s.grpc.Run(ctx))
	}()

	err := s.server.Run(ctx)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.pyroscope.Stop()
	if err != nil {
		return fmt.Errorf("failed to stop pyroscope: %w", err)
	}

	err = s.server.GracefulStop(ctx)
	if err != nil {
		return fmt.Errorf("failed to stop server: %w", err)
	}

	err = s.grpc.GracefulStop(ctx)
	if err != nil {
		return fmt.Errorf("failed to stop grpc: %w", err)
	}

	return nil
}

type config struct {
	HTTP      kgin.Config       `toml:"http"`
	Grpc      kgrpc.Config      `toml:"grpc"`
	Pyroscope kpyroscope.Config `toml:"pyroscope"`
	Etcdv3    etcdv3.Config     `toml:"etcdv3"`
}
