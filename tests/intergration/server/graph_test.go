package server

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/api/graph"
	"github.com/go-kod/kod-mono/internal/adaptor/server"
	"github.com/stretchr/testify/assert"
)

func TestGraphController(t *testing.T) {
	kod.RunTest(t, func(ctx context.Context, s server.GraphController) {
		c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
			Resolvers: s,
		})))

		var response map[string]interface{}
		c.MustPost(`{todos {id,text} }`, &response)
		assert.Len(t, response["todos"], 0)

		c.MustPost(`mutation { createTodo	(input:{text:"Fery important", userId:"121"}) { id } }`, &response)
		assert.EqualValues(t, map[string]interface{}{"createTodo": map[string]interface{}{"id": "1"}}, response)

		c.MustPost(`{todos {id,text,done} }`, &response)
		assert.EqualValues(t, map[string]interface{}{"todos": []interface{}{map[string]interface{}{"done": false, "id": "1", "text": "Fery important"}}}, response)
	}, kod.WithConfigFile("../../../config/server/dev.toml"))
}
