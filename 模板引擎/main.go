package main

import (
	"goLearning/utils"
	"os"
	"strings"
	"text/template"
)

type TplFile struct {
	FileName string
	PackageName string
	Data interface{}
}

type TplStruct struct {
	Name string `json:"name"`
	Fields []TplStructField
}

type TplStructField struct {
	Name string
	Type string
	Tag string
}


func main() {
	// 获取模板内容
	// the value at the current location in the structure
	tplDir := "/Users/ibeans/backup/workdir3/goLearning/模板引擎/"
	teplPtr, err := utils.ReadFile(tplDir + "struct.tpl")
	if err != nil {
		panic(err)
	}
	tepl := strings.TrimSpace(*teplPtr)

	// 定义模板数据
	dataStruct := TplStruct{
		Name:         "Inventory",
	}
	dataStruct.Fields = append(dataStruct.Fields, TplStructField{"Material", "string", "`json:\"material\"`"})
	dataStruct.Fields = append(dataStruct.Fields, TplStructField{"Count", "int", "`json:\"count\"`"})
	dataStruct.Fields = append(dataStruct.Fields, TplStructField{"Count1", "int", "`json:\"count\"`"})
	dataStruct.Fields = append(dataStruct.Fields, TplStructField{"Count2", "int", "`json:\"count\"`"})
	dataInfo := TplFile{
		FileName:    "result/structTpl.go",
		PackageName: "structTpl",
		Data:        dataStruct,
	}

	// 解析模板, 获取结果
	tmpl, err := template.New("test").Parse(tepl)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(tplDir + dataInfo.FileName)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, dataInfo)
	if err != nil {
		panic(err)
	}
}
