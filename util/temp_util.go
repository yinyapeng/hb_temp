package util

import (
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yinyapeng/hb_temp/bean"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func init() {
}

func GenApiTemp(data bean.TempData) {
	lower := ToLower(IgnoreCont(data.TableName, data.IgnoreTablePrefix, data.IgnoreTableSuffix))
	var outPathRoot = data.Path + "/api/" + lower
	genTemp("templ/api/interface.tpl", outPathRoot+"/"+lower+".go", "interfaceTemp", data)
	var outPathV1 = outPathRoot + "/v1"
	//增
	genTemp("templ/api/v1/add.tpl", outPathV1+"/add.go", "addTemp", data)
	//删
	genTemp("templ/api/v1/delete.tpl", outPathV1+"/delete.go", "deleteTemp", data)
	//改
	genTemp("templ/api/v1/modify.tpl", outPathV1+"/modify.go", "modifyTemp", data)
	//查单条
	genTemp("templ/api/v1/get.tpl", outPathV1+"/get.go", "getTemp", data)
	//分页查
	genTemp("templ/api/v1/list.tpl", outPathV1+"/list.go", "listTemp", data)
}

// genTemp 生成模板文件
func genTemp(localPath, outputPath, tempName string, data bean.TempData) {
	temp := gres.GetContent(localPath)
	// 创建模板
	tmpl, err := template.New(tempName).Funcs(FuncList()).Parse(string(temp))
	if err != nil {
		panic(err)
	}
	//创建输出目录
	os.MkdirAll(getDir(outputPath), os.ModePerm)
	// 创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	// 执行模板，写入输出文件
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		panic(err)
	}
}

// getDir 获取目录路径
func getDir(path string) string {
	return path[0 : len(path)-len(filepath.Base(path))]
}

func FuncList() template.FuncMap {
	return template.FuncMap{
		"ignoreCont":     IgnoreCont,
		"caseCamel":      CaseCamel,
		"caseCamelLower": CaseCamelLower,
		"toLower":        ToLower,
		"toUpper":        ToUpper,
	}
}

// IgnoreCont 处理忽略内容
func IgnoreCont(val, ignorePrefixStr, ignoreSuffixStr string) string {
	ignorePrefix := strings.Split(ignorePrefixStr, ",")
	ignoreSuffix := strings.Split(ignoreSuffixStr, ",")
	for _, prefix := range ignorePrefix {
		if strings.HasPrefix(val, prefix) {
			val = val[len(prefix):]
		}
	}
	for _, suffix := range ignoreSuffix {
		if strings.HasSuffix(val, suffix) {
			val = val[:len(val)-len(suffix)]
		}
	}
	return val
}

// CaseCamel 字符串首字母大写驼峰格式
func CaseCamel(val string) string {
	snakeStr := gstr.ToLower(val)
	return gstr.CaseCamel(snakeStr)
}

// CaseCamelLower 字符串首字母小写驼峰格式
func CaseCamelLower(val string) string {
	snakeStr := gstr.ToLower(val)
	return gstr.CaseCamelLower(snakeStr)
}

// ToLower 字符串小写
func ToLower(val string) string {
	return strings.ToLower(val)
}

// ToUpper 字符串大写
func ToUpper(val string) string {
	return strings.ToUpper(val)
}
