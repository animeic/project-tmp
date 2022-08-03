package controller

import (
	"fmt"

	"animeii.tech/request"
	"animeii.tech/response"
	"animeii.tech/service"
	"animeii.tech/utils"
	"github.com/gin-gonic/gin"
)

type SmsController struct{}

var SSC = new(SmsController)

func (ssc *SmsController) RegisterSend(c *gin.Context) {
	// 表单验证
	var form request.RegisterSendForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	fmt.Println(form)

	// 验证账号是否被注册
	as := new(service.AuthService)
	// 没找到数据err有值
	isAccount, _ := as.IsRegisterAccount(form.Account)

	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(isAccount)
	// 	response.Fail(c, response.DataNotFundErr)
	// 	return
	// }
	if isAccount {
		response.Fail(c, response.IsAccountErr)
		return
	}
	isEmail, _ := as.IsRegisterEmail(form.Email)
	// if err != nil {
	// 	fmt.Println(err)
	// 	response.Fail(c, response.DataNotFundErr)
	// 	return
	// }
	if isEmail {
		response.Fail(c, response.IsEmailErr)
		return
	}

	// 发送目标邮箱验证码
	mailTo := []string{
		form.Email,
	}
	subject := "animeii.tech"
	code := utils.MakeEmailCode()
	fmt.Println("sms_code: " + code)
	body := fmt.Sprintf(`<p style="font-size:18px;">您正在<a href="https://animeii.tech">animeii.tech</a>注册账号，验证码是：
	<span style="color:blue;font-size:24px;">%s</span></p><p style="font-size:18x;">5分钟内有效，请妥善保管。</p>`, code)
	err = new(service.SmsService).SendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, response.SendCodeErr)
		return
	}
	// 发送成功生成缓存用于验证
	// mailcode:register:animeic@163.com
	// mailcode:repass:animeic@163.com
	rd := new(service.RecDb)
	err = rd.SetEmailCode("register", mailTo[0], code)
	if err != nil {
		response.Fail(c, response.SendCodeErr)
		return
	}
	response.Success(c, response.SendCodeSucc, nil)

}
func (ssc *SmsController) RepassSend(c *gin.Context) {
	// 表单验证
	var form request.RepassSendForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}

	// 验证账号是否被注册
	as := new(service.AuthService)

	isEmail, _ := as.IsRegisterEmail(form.Email)
	// if err != nil {
	// 	response.Fail(c, response.DataNotFundErr)
	// 	return
	// }
	if !isEmail {
		response.Fail(c, response.IsNotRegisterErr)
		return
	}

	// 发送目标邮箱验证码
	mailTo := []string{
		form.Email,
	}
	subject := "animeii.tech"
	code := utils.MakeEmailCode()
	fmt.Println("sms_code: " + code)
	body := fmt.Sprintf(`<p style="font-size:18px;">您正在<a href="https://animeii.tech">animeii.tech</a>重置密码，验证码是：
	<span style="color:blue;font-size:24px;">%s</span></p><p style="font-size:18x;">5分钟内有效，请妥善保管。</p>`, code)
	err = new(service.SmsService).SendMail(mailTo, subject, body)
	if err != nil {
		response.Fail(c, response.SendCodeErr)
		return
	}
	// 发送成功生成缓存用于验证
	// mailcode:register:animeic@163.com
	// mailcode:repass:animeic@163.com
	rd := new(service.RecDb)
	err = rd.SetEmailCode("repass", mailTo[0], code)
	if err != nil {
		response.Fail(c, response.SendCodeErr)
		return
	}
	response.Success(c, response.SendCodeSucc, nil)

}
