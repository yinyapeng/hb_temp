package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yinyapeng/hb_temp/bean"
)

type GetTablesReq struct {
	g.Meta `path:"/getTables" tags:"模板生成" method:"post" summary:"获取所有表数据"`
}

type GetTablesRes struct {
	List              []bean.TreeTableData `json:"list"`
	Path              string               `json:"path"`              //文件跟路径
	IgnoreTablePrefix string               `json:"ignoreTablePrefix"` //忽略表前缀
	IgnoreTableSuffix string               `json:"ignoreTableSuffix"` //忽略表后缀
	IgnoreFieldPrefix string               `json:"ignoreFieldPrefix"` //忽略字段前缀
	IgnoreFieldSuffix string               `json:"ignoreFieldSuffix"` //忽略字段后缀
}
