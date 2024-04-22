package main

import (
	"context"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/internal/adaptor/server"
	"github.com/go-kod/kod/interceptor/kaccesslog"
	"github.com/go-kod/kod/interceptor/kmetric"
	"github.com/go-kod/kod/interceptor/krecovery"
	"github.com/go-kod/kod/interceptor/ktrace"
	"github.com/samber/lo"
)

//	@title			Swagger Example API
//	@version		2.0
//	@description	This is a sample server.
//	@termsOfService	http://swagger.io/terms/
//	@host			localhost:9527
//	@schemes		http

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	/
func main() {
	lo.Must0(kod.Run(context.Background(), func(ctx context.Context, s *server.Server) error {
		return s.Run(ctx)
	}, kod.WithInterceptors(
		krecovery.Interceptor(),
		kmetric.Interceptor(),
		ktrace.Interceptor(),
		kaccesslog.Interceptor(),
	)))
}
