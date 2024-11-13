package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yinyapeng/hb_temp/bean"
)

type GenerateReq struct {
	g.Meta            `path:"/generate" tags:"模板生成" method:"post" summary:"生成数据"`
	Path              string               `json:"path"`              //文件跟路径
	IgnoreTablePrefix string               `json:"ignoreTablePrefix"` //忽略表前缀
	IgnoreTableSuffix string               `json:"ignoreTableSuffix"` //忽略表后缀
	IgnoreFieldPrefix string               `json:"ignoreFieldPrefix"` //忽略字段前缀
	IgnoreFieldSuffix string               `json:"ignoreFieldSuffix"` //忽略字段后缀
	ForceGen          bool                 `json:"forceGen"`          //是否强制生成
	Api               bool                 `json:"api"`               //生成接口参数
	Ctrl              bool                 `json:"ctrl"`              //生成控制层
	Logic             bool                 `json:"logic"`             //生成服务层
	Dao               bool                 `json:"dao"`               //生成数据层
	Add               bool                 `json:"add"`               //添加接口
	Delete            bool                 `json:"delete"`            //删除接口
	Modify            bool                 `json:"modify"`            //修改接口
	Get               bool                 `json:"get"`               //主键获取接口
	List              bool                 `json:"list"`              //分页获取接口
	Tree              bool                 `json:"tree"`              //树形获取接口
	TableList         []bean.TreeTableData `json:"tableList"`         //生成表列表
}

type GenerateRes struct {
}
