package snowflake

import (
	"context"
	"errors"

	"github.com/bwmarrin/snowflake"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/internal/infra/redis"
)

type service struct {
	kod.Implements[Service]

	SnowflakeRepository kod.Ref[redis.SnowflakeRepository]
	node                *snowflake.Node
}

type GenReq struct {
}

type GenRes struct {
	UUID string
}

func (s *service) Init(ctx context.Context) error {
	id, err := s.SnowflakeRepository.Get().GetUniqId(ctx)
	if err != nil {
		return err
	}

	node, err := snowflake.NewNode(id)
	if err != nil {
		return err
	}
	s.node = node
	return nil
}

func (s *service) Gen(ctx context.Context, _ *GenReq) (*GenRes, error) {
	if s.node == nil {
		return nil, errors.New("snowflake node not init")
	}

	id := s.node.Generate().String()

	return &GenRes{
		UUID: id,
	}, nil
}
