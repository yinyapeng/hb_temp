package v1

import "github.com/gogf/gf/v2/frame/g"

type StaticReq struct {
	g.Meta `path:"/{type}/{name}.(js|css|png|gif)"  tags:"模板生成" method:"get" summary:"获取其他静态资源"`
}

type StaticRes struct {
}
