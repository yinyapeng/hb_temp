package util

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yinyapeng/hb_temp/bean"
	"regexp"
	"strings"
)

var (
	dataBaseData bean.DataBaseData = bean.DataBaseData{}
	ctx                            = gctx.New()
)

type databaseData struct {
	Link  string `json:"link"`
	Debug bool   `json:"debug"`
}

func init() {
	ctx := gctx.New()
	valueVal, _ := g.Cfg().Get(ctx, "database")
	swaggerPathVal, _ := g.Cfg().Get(ctx, "server.swaggerPath")
	if swaggerPathVal.String() != "" {
		if valueVal.String() != "" {
			links := make(map[string]databaseData)
			json.Unmarshal([]byte(valueVal.String()), &links)
			baseLinks := make([]bean.DataBaseLinks, len(links))
			var i = 0
			for name, item := range links {
				split := strings.Split(item.Link, ":")
				dataBaseData.DataBaseType = split[0]

				baseLinks[i].LinkName = name
				baseLinks[i].LinkCont = item.Link
				baseLinks[i].Debug = item.Debug
				baseLinks[i].Schema = getSchema(split)

				i = i + 1
			}
			dataBaseData.DataBaseLinks = baseLinks
		}
	}

}

func getSchema(links []string) string {
	switch links[0] {
	case "dm":
		compileRegex := regexp.MustCompile(`SCHEMA=([^&]+)`)
		submatch := compileRegex.FindStringSubmatch(links[3])
		if len(submatch) > 1 {
			return submatch[1]
		} else {
			return links[1]
		}
	case "mssql":
		return ""
	}
	return ""
}

func GetTables() []bean.TreeTableData {
	var tables []bean.TreeTableData
	var sql string
	switch dataBaseData.DataBaseType {
	case "dm":
		sql = `select a.TABLE_NAME Id,b.COMMENT$ Name,? PId from ALL_TABLES a 
		left join SYSTABLECOMMENTS b on a.OWNER = b.SCHNAME and a.TABLE_NAME = b.TVNAME and b.TABLE_TYPE = 'TABLE'
		where a.OWNER = ?;`
	case "mssql":
	}

	if sql != "" {
		links := dataBaseData.DataBaseLinks
		for _, item := range links {
			tables = append(tables, bean.TreeTableData{
				Id:   item.LinkName,
				Name: item.LinkName,
				PId:  "",
			})
			all, _ := g.DB(item.LinkName).GetAll(ctx, sql, item.LinkName, item.Schema)
			var rows []bean.TreeTableData
			gconv.Scan(all, &rows)
			tables = append(tables, rows...)
		}
	}

	return tables
}

func GetTabColumns(linkName, schema, tableName string) []bean.TabColumns {
	var columns []bean.TabColumns
	var sql string
	switch dataBaseData.DataBaseType {
	case "dm":
		sql = `select a.TABLE_NAME tableName,a.COLUMN_NAME columnName,a.DATA_TYPE dataType,a.DATA_LENGTH dataLength,a.NULLABLE nullable,a.COLUMN_ID columnId,b.COMMENT$ describe
		from ALL_TAB_COLUMNS a 
		left join SYSCOLUMNCOMMENTS b on a.OWNER = b.SCHNAME and a.TABLE_NAME = b.TVNAME and a.COLUMN_NAME = b.COLNAME and b.TABLE_TYPE = 'TABLE'
		where a.owner = ? AND a.TABLE_NAME=?;`
	case "mssql":
	}
	if sql != "" {
		all, _ := g.DB(linkName).GetAll(ctx, sql, schema, tableName)
		gconv.Scan(all, &columns)
	}
	return columns
}
