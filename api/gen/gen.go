package gen

import (
	"context"
	v1 "github.com/yinyapeng/hb_temp/api/gen/v1"
)

type IGenV1 interface {
	Html(ctx context.Context, req *v1.HtmlReq) (res *v1.HtmlRes, err error)
	Static(ctx context.Context, req *v1.StaticReq) (res *v1.StaticRes, err error)
	GetTables(ctx context.Context, req *v1.GetTablesReq) (res *v1.GetTablesRes, err error)
	Generate(ctx context.Context, req *v1.GenerateReq) (res *v1.GenerateRes, err error)
}
