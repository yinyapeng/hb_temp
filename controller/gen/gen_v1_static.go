package gen

import (
	"context"
	"fmt"
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
		fmt.Println("public/" + t.String() + "/" + name.String() + "." + t.String())
		temp := gres.GetContent("public/" + t.String() + "/" + name.String() + "." + t.String())
		r.Response.Header().Set("content-type", "text/"+t.String()+"; charset=utf-8")
		r.Response.Writeln(temp)
	}
	return
}
