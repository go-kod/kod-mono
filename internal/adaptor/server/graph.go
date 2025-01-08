package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/api/graph"
	"github.com/go-kod/kod-mono/internal/app/example"
	"github.com/ravilushqa/otelgqlgen"
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
	h := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: g,
	}))
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.GET{})

	h.Use(otelgqlgen.Middleware())

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
