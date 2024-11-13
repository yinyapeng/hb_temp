package gen

import (
	"context"
	"github.com/yinyapeng/hb_temp/api/gen/v1"
	"github.com/yinyapeng/hb_temp/util"
	"os"
)

func (c *ControllerV1) GetTables(ctx context.Context, req *v1.GetTablesReq) (res *v1.GetTablesRes, err error) {
	tables := util.GetTables()
	projectPath, _ := os.Getwd()
	res = &v1.GetTablesRes{
		List:              tables,
		Path:              projectPath,
		IgnoreTablePrefix: "k_,K_,d_,D_",
	}
	return
}
