package controller

import (
	"animeii.tech/dd"
	"animeii.tech/model"
	"animeii.tech/request"
	"animeii.tech/response"
	"animeii.tech/utils"
	"github.com/gin-gonic/gin"
)

type CommentController struct{}

var CTC = new(CommentController)

func (ctc *CommentController) Post(c *gin.Context) {
	// 请求参数
	var form request.CommentPost
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	// 获取操作者id
	uaid, isExist := c.Get("id")
	if !isExist {
		response.Fail(c, response.ParamsErr)
		return
	}
	uid := utils.GetUid(uaid)
	// 新增数据
	data := model.Comment{
		Uid:     uid,
		Pid:     form.Pid,
		Aid:     form.Aid,
		Content: form.Content,
	}
	var resp response.CommentPost
	result := dd.DB.Create(&data).Scan(&resp)
	if result.Error != nil {
		response.Fail(c, response.CommentPostErr)
		return
	}
	response.Success(c, response.CommentSucc, resp)
}

func (ctc *CommentController) List(c *gin.Context) {
	// aid
	// 根据aid，pid=0找到评论列表，用户列表，排序规则。最新，最热
	// 前端根据数据特征层级展示
	var form request.CommentList
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}

}
