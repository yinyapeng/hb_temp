package v1

import "github.com/gogf/gf/v2/frame/g"

type StaticReq struct {
	g.Meta `path:"/*.(html|js|css|png|gif)"  tags:"模板生成" method:"get" summary:"获取静态资源"`
}

type StaticRes struct {
}
