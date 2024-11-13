package v1

import "github.com/gogf/gf/v2/frame/g"

type HtmlReq struct {
	g.Meta `path:"/{name}.html" tags:"模板生成" method:"get" summary:"获取html静态资源"`
}

type HtmlRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
