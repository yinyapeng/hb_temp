package bean

type TreeTableData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	PId  string `json:"PId"`
}

type DataBaseData struct {
	DataBaseType  string          `json:"dataBaseType"`
	DataBaseLinks []DataBaseLinks `json:"dataBaseLinks"`
}

type DataBaseLinks struct {
	LinkName string `json:"linkName"`
	LinkCont string `json:"linkCont"`
	Debug    bool   `json:"debug"`
	Schema   string `json:"schema"`
}

// TempData 模板生成数据
type TempData struct {
	Path              string       `json:"path"`              //文件跟路径
	LinkName          string       `json:"linkName"`          //链接名称
	TableName         string       `json:"tableName"`         //表名称
	TableDesc         string       `json:"tableDesc"`         //表描述
	IgnoreTablePrefix string       `json:"ignoreTablePrefix"` //忽略表前缀
	IgnoreTableSuffix string       `json:"ignoreTableSuffix"` //忽略表后缀
	IgnoreFieldPrefix string       `json:"ignoreFieldPrefix"` //忽略字段前缀
	IgnoreFieldSuffix string       `json:"ignoreFieldSuffix"` //忽略字段后缀
	TabColumns        []TabColumns `json:"tabColumns"`        //表字段列表

}

type TabColumns struct {
	TableName  string `json:"tableName"`
	ColumnName string `json:"columnName"`
	DataType   string `json:"dataType"`
	DataLength int    `json:"dataLength"`
	ColumnId   int    `json:"columnId"`
	Describe   string `json:"describe"`
}
