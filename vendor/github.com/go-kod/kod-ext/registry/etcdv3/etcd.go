package etcdv3

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"dario.cat/mergo"
	"github.com/go-kod/kod"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	gresolver "google.golang.org/grpc/resolver"
)

// nolint
type Config struct {
	Endpoints []string
	Timeout   time.Duration
	TTL       int
}

// nolint

// nolint
func (r Config) Build(ctx context.Context) (*client, error) {

	err := mergo.Merge(&r, Config{
		Timeout: 3 * time.Second,
		TTL:     60,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to merge config: %w", err)
	}

	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   r.Endpoints,
		DialTimeout: r.Timeout,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %w", err)
	}

	cc := new(client)
	cc.client = etcd

	if r.TTL > 0 {

		ss, err := concurrency.NewSession(cc.client, concurrency.WithTTL(r.TTL), concurrency.WithContext(ctx))
		if err != nil {
			return nil, fmt.Errorf("failed to create etcd session: %w", err)
		}

		cc.session = ss
	}

	manager, err := endpoints.NewManager(cc.client, cc.registryPrefix(ctx))
	if err != nil {
		return cc, err
	}

	cc.manager = manager

	return cc, nil
}

type client struct {
	client  *clientv3.Client
	manager endpoints.Manager
	session *concurrency.Session
}

// nolint
func (r *client) Register(ctx context.Context, scheme, addr string) error {
	opts := []clientv3.OpOption{}
	if r.session != nil {
		opts = append(opts, clientv3.WithLease(r.session.Lease()))
	}

	slog.InfoContext(ctx, "Register service", "key", r.registryKey(ctx, scheme, addr))

	err := r.manager.AddEndpoint(context.Background(), r.registryKey(ctx, scheme, addr), endpoints.Endpoint{
		Addr: addr,
	}, opts...)
	if err != nil {
		return err
	}

	return nil
}

// nolint
func (r *client) UnRegister(ctx context.Context, scheme, addr string) error {
	slog.InfoContext(ctx, "UnRegister service", "key", r.registryKey(ctx, scheme, addr))

	return r.manager.DeleteEndpoint(context.Background(), r.registryKey(ctx, scheme, addr))
}

// nolint
func (r *client) registryPrefix(ctx context.Context) string {
	return fmt.Sprintf("%s/%s", kod.FromContext(ctx).Config().Env, kod.FromContext(ctx).Config().Name)
}

// nolint
func (r *client) registryKey(ctx context.Context, scheme, addr string) string {
	return fmt.Sprintf("%s/%s/%s/%s", kod.FromContext(ctx).Config().Env, kod.FromContext(ctx).Config().Name, scheme, addr)
}

// nolint
func (s *client) NewBuilder(ctx context.Context) (gresolver.Builder, error) {
	return builder{c: s.client}, nil
}
