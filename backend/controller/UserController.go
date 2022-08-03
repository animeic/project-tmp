package controller

import (
	"fmt"

	"animeii.tech/request"
	"animeii.tech/response"
	"animeii.tech/service"
	"animeii.tech/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var USC = new(UserController)

// new 获取用户简易信息
func (usc *UserController) Info(c *gin.Context) {
	// 验证表单
	var form request.UserInfoForm
	err := c.ShouldBindJSON(&form)
	// 表单规则不通过
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	// 获取用户信息
	us := new(service.UserService)
	res, err := us.GetInfo(form.Id)
	// 查询不到用户
	if err != nil {
		response.Fail(c, response.DataNotFundErr)
		return
	}
	// 返回用户信息
	response.Success(c, response.GetUserInfoSucc, res)
}

// new 获取用户简易信息
func (usc *UserController) InfoNow(c *gin.Context) {
	any_uid, _ := c.Get("id")
	fmt.Println(any_uid)
	uid := utils.GetUid(any_uid)
	fmt.Println(uid)
	if uid < 1 {
		response.Fail(c, response.DataNotFundErr)
		return
	}
	// 获取用户信息
	us := new(service.UserService)
	res, err := us.GetInfo(uid)
	// 查询不到用户
	if err != nil {
		response.Fail(c, response.DataNotFundErr)
		return
	}
	// 返回用户信息
	response.Success(c, response.GetUserInfoSucc, res)
}
