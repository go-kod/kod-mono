package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/api/graph"
	"github.com/go-kod/kod-mono/internal/app/example"
)

type resolver struct {
	kod.Implements[GraphController]
	g kod.Ref[example.GraphService]
}

func (r *resolver) Query() graph.QueryResolver {
	return r.g.Get()
}

func (r *resolver) Mutation() graph.MutationResolver {
	return r.g.Get()
}

func graphqlHandler(g GraphController) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: g,
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
