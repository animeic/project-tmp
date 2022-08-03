package router

import (
	"animeii.tech/controller"
	"animeii.tech/middleware"
	"animeii.tech/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	// api 跨域必须放在路由前面
	app.Use(middleware.Cors(), middleware.CustomRecovery())

	upload := app.Group("/upload")
	upload.POST("/image", controller.UD.Image)

	article := app.Group("/article")
	article.POST("/detail", controller.ARC.Detail)
	article.POST("/list", controller.ARC.List)

	admin := app.Group("admin")
	admin.Use(middleware.Auth(service.AppGuardName)) // 验证token
	admin.POST("/article/post", controller.ARC.Post)
	admin.POST("/article/editDetail", controller.ARC.EditDetail) // 用于编辑回显
	admin.POST("/article/edit", controller.ARC.Edit)             // 用于编辑回显

	user := app.Group("/user")
	// user.Use(middleware.Auth(service.AppGuardName))
	user.POST("/info", controller.USC.Info)
	user.Use(middleware.Auth(service.AppGuardName)).POST("/infoNow", controller.USC.InfoNow)

	auth := app.Group("/auth")
	// 暂时禁用，仅供自己使用
	//auth.POST("/register", controller.AC.Register)
	auth.POST("/login", controller.AC.Login)
	auth.POST("/repass", controller.AC.Repass)

	sms := app.Group("/sms")
	sms.POST("/registerCode", controller.SSC.RegisterSend)
	sms.POST("/repassCode", controller.SSC.RepassSend)

}
