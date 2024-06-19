package server

import (
	"net/http"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-mono/internal/app/example"
	kgin "github.com/go-kod/kod/ext/server/kgin"
	"google.golang.org/grpc/status"
)

type ginImpl struct {
	kod.Implements[GinController]

	comp kod.Ref[example.Service]
}

func registerHTTP(s *kgin.Server, c GinController, graph GraphController) {
	s.GET("/uniqueId", c.UniqueID)

	s.Handle("POST", "/graphql", graphqlHandler(graph))
	s.Handle("GET", "/", playgroundHandler())
}

// @Summary		Get a Unique ID
// @Description	get unique ID
// @Consume		x-www-form-urlencoded
// @Produce		json
// @Param			request	query		example.TestReq	true	"请求参数"
// @Success		200		{object}	example.TestRes	"ok"
// @Router			/uniqueId [get]
func (c *ginImpl) UniqueID(ctx *kgin.Context) {
	req := &example.TestReq{
		Name: "",
	}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, status.Convert(err).Proto())
		return
	}

	res, err := c.comp.Get().UniqueID(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, status.Convert(err).Proto())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
