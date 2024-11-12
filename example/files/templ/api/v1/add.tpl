package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddReq struct {
	g.Meta   `path:"/{{ (ignoreCont .TableName .IgnoreTablePrefix .IgnoreTableSuffix) | toLower }}/add" tags:"{{ .TableDesc }}服务" method:"post" summary:"{{ .TableDesc }}添加"`
    {{range $key, $item := .TabColumns}}
    {{ $item.ColumnName | caseCamel }}
    {{end}}
}
type AddRes struct {
	代码 string `json:"代码" dc:"{{ .TableDesc }}ID"`
}
