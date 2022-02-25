package main

import (
	"archive/zip"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	//基础数据
	userID := int64(30000007)
	dirPath := fmt.Sprintf("./csv/玩家_%d_历史记录_%d", userID, time.Now().Unix())
	zipPath := fmt.Sprintf("./csv/玩家_%d_历史记录_%d.zip", userID, time.Now().Unix())

	//数据准备
	title := []string{"序号", "国家", "ID"}
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}

	DoExportData(userID, "充值记录", dirPath, title, data)
	DoExportData(userID, "登陆记录", dirPath, title, data)
	DoExportData(userID, "道具变化", dirPath, title, data)
	Zip(dirPath, zipPath)
}

/*
DoExportData
@Desc 	创建文件-写入文件
@Param 	userID 		int64		玩家ID
@Param 	excelName 	string		excel名称
@Param 	dir 		string		文件夹地址
@Param 	title 		[]string	内容标题
@Param 	data 		[][]string	内容数据
*/
func DoExportData(userID int64, excelName string, dir string, title []string, data [][]string) (err error) {
	//创建文件夹
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Printf("MkdirAll is faile. userID:%v, err:%v", userID, err.Error())
		return
	}

	//创建文件
	fileName := fmt.Sprintf("%s/%s.xls", dir, excelName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("NewFile is faile. userID:%v, err:%v", userID, err.Error())
		return
	}
	defer file.Close()

	//写入文件
	err = Write(file, title, data)
	if err != nil {
		fmt.Printf("Write is faile. userID:%v, err:%v", userID, err.Error())
		return
	}

	//success
	return
}

/*
Write
@Desc 	写入数据到excel
@Param	file 	*os.File	文件信息
@Param	title 	[]string	文件标题
@Param	content [][]string	文件内容
*/
func Write(file *os.File, title []string, content [][]string) (err error) {
	if file == nil {
		return errors.New("file is nil")
	}

	//写入UTF-8 BOM
	_, err = file.WriteString("\xEF\xBB\xBF")
	if err != nil {
		return
	}

	w := csv.NewWriter(file)

	//写入标题
	err = w.Write(title)
	if err != nil {
		return
	}

	//写入内容
	err = w.WriteAll(content)
	if err != nil {
		return
	}

	w.Flush()
	return nil
}

/*
Zip
@Desc 	压缩文件夹
@Param 	srcDir 		string	文件夹地址
@Param 	zipFileName string	压缩包名称
*/
func Zip(srcDir string, zipFileName string) {
	//预防：旧文件无法覆盖
	os.RemoveAll(zipFileName)

	//创建：zip文件
	zipfile, _ := os.Create(zipFileName)
	defer zipfile.Close()

	//打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	//遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		//如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		//获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`\`)

		//判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			//设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		//创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}

		//success
		return nil
	})
}

/*
rowsToArrayAndMap
@Desc 	SQL数据转成数组或者map
@Return listArr [][]string					为了存csv
@Return listMap []map[string]interface{}	原先的保留住
*/
func rowsToArrayAndMap(rows *sql.Rows) (listArr [][]string, listMap []map[string]interface{}) {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}

	for rows.Next() {
		_ = rows.Scan(cache...)
		itemArr := make([]string, columnLength, columnLength)
		itemMap := make(map[string]interface{})
		for key, data := range cache {
			value := *data.(*interface{})
			itemArr[key] = fmt.Sprintf("%s", value)
			itemMap[columns[key]] = value
		}
		listArr = append(listArr, itemArr)
		listMap = append(listMap, itemMap)
	}
	_ = rows.Close()
	return
}
