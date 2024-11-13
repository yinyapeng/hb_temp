package gen

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/yinyapeng/hb_temp/api/gen/v1"
)

func (c *ControllerV1) Static(ctx context.Context, req *v1.StaticReq) (res *v1.StaticRes, err error) {
	r := g.RequestFromCtx(ctx)
	if SwaggerPath == "" {
		r.Response.Writeln("Not Found")
	} else {
		t := r.Get("type")
		name := r.Get("name")
		temp := gres.GetContent(t.String() + "/" + name.String() + "." + t.String())
		r.Response.Writeln(temp)
	}
	return
}
