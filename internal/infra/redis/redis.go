package redis

import (
	"context"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-ext/client/kredis"
)

type snowflake struct {
	kod.Implements[SnowflakeRepository]
	kod.WithConfig[kredis.Config]

	redis *kredis.Client
}

func (s *snowflake) Init(context.Context) error {
	s.redis = s.Config().Build()
	return nil
}

func (s *snowflake) GetUniqId(ctx context.Context) (int64, error) {
	return s.redis.Incr(ctx, "uniq_id").Result()
}
