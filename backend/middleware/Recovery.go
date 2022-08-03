package middleware

import (
	"animeii.tech/response"

	"animeii.tech/dd"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   dd.Cf.Log.RootDir + "/" + dd.Cf.Log.Filename,
			MaxSize:    dd.Cf.Log.MaxSize,
			MaxBackups: dd.Cf.Log.MaxBackups,
			MaxAge:     dd.Cf.Log.MaxAge,
			Compress:   dd.Cf.Log.Compress,
		},
		response.ServerError)
}
