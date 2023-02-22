package main

import (
	"encoding/base64"
	"fmt"
	"git.jiaxianghudong.com/go/logs"
	"github.com/tealeg/xlsx"
	"strings"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var configJX Config
var configWL Config
var configTEST Config

//Config 配置
type Config struct {
	filePath string
	mysql    *MysqlOption
}

//Data excel数据
type Data struct {
	ID      int64
	Code    string
	OldCode string
}

// MysqlOption mysql 配置
type MysqlOption struct {
	Driver  string
	MaxOpen int
	MaxIdle int
}

func init() {

}

func main() {
	//选择更新
	config := configJX

	//连接数据库
	db, err := gorm.Open(mysql.Open(config.mysql.Driver), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//读取信息
	res, err := ReadFile(config.filePath)
	if err != nil || len(res) <= 0 {
		fmt.Println("读取文件有错误")
		return
	}

	//解码校验
	decode := decodeCode(res)
	if len(res) != len(decode) {
		fmt.Println("解码有错误")
		return
	}
	return
	//连接数据库
	wrongCount := 0
	successCount := 0
	total := len(decode)
	for _, v := range decode {
		update := make(map[string]interface{})
		update["code"] = v.Code
		where := make(map[string]interface{})
		where["id"] = v.ID
		where["code"] = v.OldCode
		err = db.Table("exchange_code_record").Where(where).Updates(update).Error
		if err != nil {
			logs.Errorf("更新失败:%+v", v)
			wrongCount++
			continue
		}
		successCount++
		logs.Infof("更新中~ 总：%v, 成功：%v, 失败：%v", total, successCount, wrongCount)
	}

	logs.Infof("更新结束，总：%v, 成功：%v, 失败：%v", total, successCount, wrongCount)
	select {}
	return
}

/*
decodeCode
@Author LuoQiang 2022-12-22 15:34
@Desc 	解码所有兑换码
@Param 	res []Data	所有待解码的信息
*/
func decodeCode(res []Data) (decode []Data) {
	decode = make([]Data, len(res))
	for k, v := range res {
		//尝试解码
		code, ok := CheckCode(v.Code)
		if !ok {
			continue
		}

		decode[k] = Data{
			ID:      v.ID,
			Code:    code,
			OldCode: v.Code,
		}
	}
	return
}

/*
CheckCode
@Author LuoQiang 2022-12-22 15:34
@Desc 	校验兑换码
@Param 	code string	兑换码信息
*/
func CheckCode(code string) (str string, res bool) {
	//截取
	strArr := strings.Split(code, "_")
	if len(strArr) < 2 {
		return
	}

	//前缀不匹配
	if strArr[0] != "BYDRAWARD" {
		return
	}

	//des解码
	key := "jxwlbydr"
	strDecrypted, ok := DesDecryptStr(strArr[1], key)
	if ok != "ok" {
		return
	}

	//base64解码
	dst, err := base64.StdEncoding.DecodeString(strDecrypted)
	if err != nil {
		return
	}

	//success
	res = true
	str = string(dst)
	return
}

/*
ReadFile
@Author LuoQiang 2022-12-22 15:34
@Desc 	读取excel信息
@Param 	filePath string		文件路径
*/
func ReadFile(filePath string) (res []Data, err error) {
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	//只读取第一页的数据
	sheet1 := file.Sheets[0]
	res = make([]Data, 0)
	for _, row := range sheet1.Rows {
		if len(row.Cells) < 2 {
			continue
		}

		//读取ID
		id, err := row.Cells[0].Int64()
		if err != nil || id <= 0 {
			continue
		}

		//读取code
		code := row.Cells[1].String()
		res = append(res, Data{
			ID:   id,
			Code: code,
		})
	}

	return
}
