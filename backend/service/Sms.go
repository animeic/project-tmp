package service

import (
	"context"
	"fmt"
	"strconv"

	"animeii.tech/dd"
	"gopkg.in/gomail.v2"
)

type SmsService struct{}

const (
	USER string = "animeic_vps@163.com"
	PASS string = "KCEWJSRYLYOVAOBD"
	HOST string = "smtp.163.com"
	PORT string = "465"
)

func (as *SmsService) SendMail(mailTo []string, subject string, body string) error {
	mailConfig := map[string]string{
		"user": USER,
		"pass": PASS,
		"host": HOST,
		"port": PORT,
	}
	port, _ := strconv.Atoi(mailConfig["port"])
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConfig["user"], "animeii官方"))
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConfig["host"], port, mailConfig["user"], mailConfig["pass"])

	err := d.DialAndSend(m)
	return err
}

// 验证码是否正确
func (ssc *SmsService) IsCode(etype string, toemail string, ecode string) (is bool, err error) {
	rdc := dd.Redis
	cxt := context.Background()
	key := fmt.Sprintf("ecode:%s:%s", etype, toemail)
	code1, err1 := rdc.Get(cxt, key).Result()
	if err1 != nil {
		err = err1
	}

	if code1 == ecode {
		is = true
	}
	return
}
