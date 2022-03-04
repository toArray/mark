package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {
	sender := "123456789@qq.com" //发送者qq邮箱
	authCode := "auth_code"      //qq邮箱授权码
	mailTitle := "邮件标题"          //邮件标题
	mailBody := "邮件内容"           //邮件内容,可以是html

	//接收者邮箱列表
	mailTo := []string{
		"11111111@qq.com",
		"22222222@qq.com",
		"33333333@qq.com",
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)       //发送者腾讯企业邮箱账号
	m.SetHeader("To", mailTo...)      //接收者邮箱列表
	m.SetHeader("Subject", mailTitle) //邮件标题
	m.SetBody("text/html", mailBody)  //邮件内容,可以是html

	//添加附件
	zipPath := "./user/temp.zip"
	m.Attach(zipPath)

	//发送邮件服务器、端口、发送者qq邮箱、qq邮箱授权码
	//服务器地址和端口是腾讯的
	d := gomail.NewDialer("smtp.qq.com", 587, sender, authCode)
	if err := d.DialAndSend(m); err != nil {
		log.Println("send mail failed", err)
		return
	}

	log.Println("success")
}
