package example

import (
	"context"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/api/graph"
	"github.com/go-kod/kod-mono/api/graph/model"
	"github.com/spf13/cast"
)

type graphImpl struct {
	kod.Implements[GraphService]

	todos []*model.Todo
}

var (
	_ graph.QueryResolver    = &graphImpl{}
	_ graph.MutationResolver = &graphImpl{}
)

func (s *graphImpl) Todos(ctx context.Context) ([]*model.Todo, error) {
	return s.todos, nil
}

func (s *graphImpl) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	data := &model.Todo{
		ID:   cast.ToString(int64(len(s.todos) + 1)),
		Text: input.Text,
		Done: false,
	}

	s.todos = append(s.todos, data)

	return data, nil
}
