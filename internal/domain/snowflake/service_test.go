package snowflake

import (
	"context"
	"testing"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/internal/infra/redis"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSnowflake(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {

		kod.RunTest(t, func(ctx context.Context, s Service) {
			res, err := s.Gen(ctx, &GenReq{})
			assert.Nil(t, err)
			assert.NotEmpty(t, res.UUID)
		}, kod.WithConfigFile("../../../config/kod-dev.toml"))
	})

	t.Run("mock", func(t *testing.T) {
		fake := redis.NewMockSnowflakeRepository(gomock.NewController(t))
		fake.EXPECT().GetUniqId(gomock.Any()).Return(int64(1), nil)

		kod.RunTest(t, func(ctx context.Context, s Service) {
			res, err := s.Gen(ctx, &GenReq{})
			assert.Nil(t, err)
			assert.NotEmpty(t, res.UUID)
		}, kod.WithFakes(kod.Fake[redis.SnowflakeRepository](fake)))
	})

}

func BenchmarkSnowflake(b *testing.B) {
	b.Run("snowflake", func(b *testing.B) {

		kod.RunTest(b, func(ctx context.Context, s Service) {
			res, err := s.Gen(ctx, &GenReq{})
			assert.Nil(b, err)
			assert.NotEmpty(b, res.UUID)
		}, kod.WithConfigFile("../../../config/kod-dev.toml"))
	})
}
