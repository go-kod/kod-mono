//go:build tools
// +build tools

package kod

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/cosmtrek/air"
	_ "github.com/go-kod/kod/cmd/kod"
	_ "github.com/swaggo/swag/cmd/swag"
	_ "go.uber.org/mock/mockgen"
)
