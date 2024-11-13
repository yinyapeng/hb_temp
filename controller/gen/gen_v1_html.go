package gen

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/yinyapeng/hb_temp/api/gen/v1"
)

func (c *ControllerV1) Html(ctx context.Context, req *v1.HtmlReq) (res *v1.HtmlRes, err error) {
	r := g.RequestFromCtx(ctx)
	if SwaggerPath == "" {
		r.Response.Writeln("Not Found")
	} else {
		name := r.Get("name")
		temp := gres.GetContent("public/" + name.String() + ".html")
		r.Response.Writeln(temp)
	}
	return
}
