package gen

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/yinyapeng/hb_temp/api/gen/v1"
	"strings"
)

func (c *ControllerV1) Static(ctx context.Context, req *v1.StaticReq) (res *v1.StaticRes, err error) {
	r := g.RequestFromCtx(ctx)
	if SwaggerPath == "" {
		r.Response.Header().Set("content-type", "text/html; charset=utf-8")
		r.Response.Writeln("Not Found")
	} else {
		router := r.GetRouter(".(html|js|css|png|gif)").String()

		suffix := router[strings.LastIndex(router, ".")+1:]
		r.Response.Header().Set("content-type", "text/"+suffix+"; charset=utf-8")

		temp := gres.GetContent("public/" + router)
		r.Response.Writeln(temp)
	}
	return
}
