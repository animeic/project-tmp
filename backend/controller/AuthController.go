package controller

import (
	"fmt"

	"animeii.tech/model"
	"animeii.tech/request"
	"animeii.tech/response"
	"animeii.tech/service"
	"animeii.tech/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var AC = new(AuthController)

func (ac *AuthController) Register(c *gin.Context) {
	// 表单验证
	var form request.RegisterForm
	err := c.ShouldBindJSON(&form)
	// 表单验证失败
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return

	}

	ssc := new(service.SmsService)
	// 验证验证码是否正
	isCode, err := ssc.IsCode("register", form.Email, form.Ecode)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, response.DataNotFundErr)
		return
	}
	if !isCode {
		response.Fail(c, response.EmailCodeErr)
		return
	}

	// 注册用户
	user := &model.User{
		Account:  form.Account,
		Email:    form.Email,
		Password: utils.Sha256(form.Password),
	}
	as := new(service.AuthService)
	userinfo, err := as.Register(user)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, response.DataNotFundErr)
		return
	}

	// 注册成功返回token
	jwt := new(service.JwtService)
	tokenData, _, err := jwt.CreateToken(service.AppGuardName, userinfo)
	if err != nil {
		response.Fail(c, response.GenTokenErr)
		return
	}

	response.Success(c, response.RegisterSucc, response.Auth{UserInfo: userinfo, Token: tokenData})
}

// 根据用户名登录
func (ac *AuthController) Login(c *gin.Context) {
	// 接收用户名和密码
	var form request.LoginForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}

	// 验证用户名或者密码是否注册
	as := new(service.AuthService)
	userinfo, err := as.CheckAccount(&form)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, response.DataNotFundErr)
		return
	}
	fmt.Printf("auth-login: %+v\n", userinfo)

	if userinfo.Id < 1 {
		response.Fail(c, response.CheckAccountErr)
		return
	}

	// 登录生成token
	jwt := new(service.JwtService)
	tokenData, _, err := jwt.CreateToken(service.AppGuardName, userinfo)
	if err != nil {
		response.Fail(c, response.GenTokenErr)
		return
	}
	response.Success(c, response.LoginSucc, response.Auth{UserInfo: userinfo, Token: tokenData})
}

// 重置密码
func (uc *AuthController) Repass(c *gin.Context) {
	var form request.RepassForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	// 验证验证码
	ssc := new(service.SmsService)
	isCode, err := ssc.IsCode("repass", form.Email, form.Ecode)
	if err != nil {
		response.Fail(c, response.EmailCodeErr)
		return
	}
	if !isCode {
		response.Fail(c, response.EmailCodeErr)
		return
	}
	// 重置密码
	user_info, err := new(service.AuthService).Repss(&form)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, response.RepassErr)
		return
	}
	if user_info.Id < 1 {
		response.Fail(c, response.RepassErr)
		return

	}
	jwt := new(service.JwtService)
	tokenData, _, err := jwt.CreateToken(service.AppGuardName, user_info)
	if err != nil {
		response.Fail(c, response.GenTokenErr)
		return
	}
	response.Success(c, response.RepassSucc, response.Auth{UserInfo: user_info, Token: tokenData})
}
