package gen

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	SwaggerPath string
)

func init() {
	ctx := gctx.New()
	swaggerPathVal, _ := g.Cfg().Get(ctx, "server.swaggerPath")
	SwaggerPath = swaggerPathVal.String()
}
