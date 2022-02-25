package main

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"time"
)

type A struct {
	ADADAS   string
	BVVASDAS int32
}

func main() {
	//创建文件
	file := xlsx.NewFile()

	//添加文件内容
	title := []interface{}{"姓名", "性别", "	年龄"}
	content := make([][]interface{}, 0)
	content = append(content, []interface{}{"姓名", "性别", A{
		ADADAS:   "3213",
		BVVASDAS: 123312,
	}, 31231412, "婚配", "现居地"})
	err := AddSheetContent(file, "登陆记录", title, content)
	err = AddSheetContent(file, "充值记录", title, content)
	err = AddSheetContent(file, "进出房间记录", title, content)
	err = AddSheetContent(file, "道具变化记录", title, content)

	//保存文件
	fileName := fmt.Sprintf("玩家_%d_历史记录_%d.xlsx", 30000007, time.Now().Unix())
	path := fmt.Sprintf("%s%s", "./Excel/", fileName)
	err = file.Save(path)
	if err != nil {
		fmt.Printf("file save is failed. fileNmae:%s, err:%v", fileName, err)
		return
	}
}

/*
AddSheetContent
@Desc 	添加一页sheet内容
@Param  sheetName 	string		页名称
@Param  title 		[]string	标题
@Param  content 	[]string	内容
*/
func AddSheetContent(file *xlsx.File, sheetName string, title []interface{}, content [][]interface{}) (err error) {
	//创建sheet
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return
	}

	//设置标题
	titleRow := sheet.AddRow()
	xlsRow := NewRow(titleRow, title)
	err = xlsRow.SetRowTitle()
	if err != nil {
		return
	}

	//添加内容
	for _, value := range content {
		currentRow := sheet.AddRow()
		contentXlsRow := NewRow(currentRow, value)
		err = contentXlsRow.GenerateRow()
		if err != nil {
			return
		}
	}

	return nil
}

type XlsxRow struct {
	Row  *xlsx.Row
	Data []interface{}
}

func NewRow(row *xlsx.Row, data []interface{}) *XlsxRow {
	return &XlsxRow{
		Row:  row,
		Data: data,
	}
}

/*
SetRowTitle
@Desc	设置行title
*/
func (row *XlsxRow) SetRowTitle() error {
	return generateRow(row.Row, row.Data)
}

/*
GenerateRow
@Desc	创建一行
*/
func (row *XlsxRow) GenerateRow() error {
	return generateRow(row.Row, row.Data)
}

/*
generateRow
@Desc 创建行并填充数据
*/
func generateRow(row *xlsx.Row, rowStr []interface{}) error {
	if rowStr == nil {
		return errors.New("rowStr count is zero")
	}
	for _, v := range rowStr {
		cell := row.AddCell()
		cell.SetValue(v)
	}
	return nil
}
