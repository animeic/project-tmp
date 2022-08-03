package controller

import (
	"animeii.tech/dd"
	"animeii.tech/model"
	"animeii.tech/request"
	"animeii.tech/response"
	"animeii.tech/service"
	"animeii.tech/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ArticleController struct{}

var ARC = new(ArticleController)

func (arc *ArticleController) Post(c *gin.Context) {

	// 表单数据结构 50rune
	var form request.ArticlePost
	var dduid uint

	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	uid, _ := c.Get("id")
	dduid = utils.GetUid(uid)

	title, err := utils.ParseMd(form.Content, "# ")
	if err != nil {
		response.Fail(c, response.ParseTextErr)
		return
	}
	// title := "fdjflkdj"
	summary, err := utils.ParseMd(form.Content, "> ")
	if err != nil {
		response.Fail(c, response.ParseTextErr)
		return
	}

	data := model.Article{
		Uid:     dduid,
		Content: form.Content,
		Html:    form.Html,
		Title:   title,
		Summary: summary,
	}
	// 返回新增数据
	var article response.ArticlePost
	res := dd.DB.Create(&data).Scan(&article)
	if res.Error != nil {
		// 这里是数据未变动发生
		response.Fail(c, response.TextNotChangeErr)
		return
	}
	response.Success(c, response.ArticlePostSucc, article)
}

func (arc *ArticleController) Edit(c *gin.Context) {
	var form request.ArticleEdit
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}

	// 删除已有缓存
	rdcKey := fmt.Sprintf("article:%d", form.Id)
	dd.Redis.Del(context.Background(), rdcKey)

	title, err := utils.ParseMd(form.Content, "# ")
	if err != nil {
		response.Fail(c, response.ParseTextErr)
		return
	}
	summary, err := utils.ParseMd(form.Content, "> ")
	if err != nil {
		response.Fail(c, response.ParseTextErr)
		return
	}

	data := model.Article{
		Content: form.Content,
		Html:    form.Html,
		Title:   title,
		Summary: summary,
	}
	// 返回新增数据
	var retid response.ArticlePost
	res := dd.DB.Model(&model.Article{}).Where("id = ?", form.Id).Updates(data).Scan(&retid)
	if res.Error != nil {
		// 这里是数据未变动发生
		response.Fail(c, response.UpdataDataErr)
		return
	}

	// 更新缓存
	var articleDetail response.ArticleDetail
	dd.DB.First(&model.Article{}, form.Id).Scan(&articleDetail)
	us := new(service.UserService)
	userInfo, _ := us.GetInfo(retid.Uid)
	articleDetail.UserInfo = userInfo
	dd.Redis.SetEX(context.Background(), rdcKey, &articleDetail, 5*time.Minute)

	response.Success(c, response.ArticlePostSucc, retid)
}

func (arc *ArticleController) EditDetail(c *gin.Context) {
	// id 编辑请求回显
	var form request.ArticleDetail
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	var res response.AtEditDetail
	// 查询对应id数据
	result := dd.DB.Select("id", "uid", "content").First(&model.Article{}, "id = ?", form.Id).Scan(&res)
	if result.Error != nil {
		response.Fail(c, response.DataNotFundErr)
		return
	}

	response.Success(c, response.GetDataSucc, res)
}

func (arc *ArticleController) List(c *gin.Context) {
	var form request.ArticleList
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	var lists []response.AtInfo  // 返回数据列表，根据需求过滤字段后
	var articles []model.Article // 查询数据集需要切片
	offset := (form.CurPage - 1) * form.PageSize
	result := dd.DB.Order("id desc").Limit(int(form.PageSize)).Offset(int(offset)).Find(&articles).Scan(&lists)
	if result.RowsAffected > 0 {
		us := new(service.UserService)
		for i := 0; i < len(lists); i++ {
			uid := lists[i].Uid
			user_info, _ := us.GetInfo(uid)
			lists[i].UserInfo = user_info
		}
	} else {
		response.Fail(c, response.DataNotFundErr)
		return

	}
	if result.Error != nil {
		response.Fail(c, response.DataNotFundErr)
		return
	}

	var count int64
	dd.DB.Select("id").Find(&articles).Count(&count)
	ret := response.ArticleList{
		Lists: lists,
		Total: uint(count),
	}

	response.Success(c, response.GetListSucc, ret)
}

func (arc *ArticleController) Detail(c *gin.Context) {
	// 表单验证
	var form request.ArticleDetail
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	// 读取mysql数据库
	var article model.Article // 哪张表

	dd.DB.Model(&article).Where("id = ?", form.Id).Update("views", gorm.Expr("views + ?", 1))

	var res response.ArticleDetail // 返回
	// 读取redis中数据 字符串转化为结构体返回
	redKey := fmt.Sprintf("article:%d", form.Id)
	err1 := dd.Redis.Get(context.Background(), redKey).Scan(&res)

	if err1 == nil {
		response.Success(c, response.ArticleGetSucc, res)
		return
	}

	// 查询文章
	result := dd.DB.First(&article, &form).Scan(&res)
	if result.Error != nil {
		response.Fail(c, response.DataNotFundErr)
		return
	}

	// 获取用户信息
	uid := res.Uid
	us := new(service.UserService)
	userInfo, _ := us.GetInfo(uid)
	res.UserInfo = userInfo
	dd.Redis.SetEX(context.Background(), redKey, &res, 5*time.Minute)
	response.Success(c, response.ArticleGetSucc, res)

}
