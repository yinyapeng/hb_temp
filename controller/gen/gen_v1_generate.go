package gen

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yinyapeng/hb_temp/api/gen/v1"
)

func (c *ControllerV1) Generate(ctx context.Context, req *v1.GenerateReq) (res *v1.GenerateRes, err error) {
	g.Dump(req)
	return
}
