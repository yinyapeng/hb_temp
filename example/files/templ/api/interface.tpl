package {{ (ignoreCont .TableName .IgnoreTablePrefix .IgnoreTableSuffix) | toLower }}

import (
	"context"
)

type I{{ (ignoreCont .TableName .IgnoreTablePrefix .IgnoreTableSuffix) | caseCamel }}V1 interface {
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Modify(ctx context.Context, req *v1.ModifyReq) (res *v1.ModifyRes, err error)
}
