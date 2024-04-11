package gin

import (
	"net/http"

	"github.com/go-kod/kod"
	gin "github.com/go-kod/kod/ext/server/kgin"
	kgin "github.com/go-kod/kod/ext/server/kgin"
	"github.com/go-kod/kod-mono/docs"
	"github.com/go-kod/kod-mono/internal/app/example"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc/status"
)

type controller struct {
	kod.Implements[Controller]

	comp kod.Ref[example.Service]
}

func Register(s *kgin.Server, c Controller) {
	docs.SwaggerInfo.BasePath = "/"

	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	s.GET("/uniqueId", c.UniqueID)
}

// @Summary		Get a Unique ID
// @Description	get unique ID
// @Consume		x-www-form-urlencoded
// @Produce		json
// @Param			request	query		example.TestReq	true	"请求参数"
// @Success		200		{object}	example.TestRes	"ok"
// @Router			/uniqueId [get]
func (c *controller) UniqueID(ctx *gin.Context) {
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
