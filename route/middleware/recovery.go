package middleware

import (
	"WebMsg/global"
	"WebMsg/utils/game"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	//加载配置
	conf := global.App.Config
	timeStr := time.Now().Format("200601/02")
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   conf.Log.RootDir + "/" + timeStr + "_err.log",
			MaxSize:    conf.Log.MaxSize,
			MaxBackups: conf.Log.MaxBackups,
			MaxAge:     conf.Log.MaxAge,
			Compress:   conf.Log.Compress,
		},
		game.ServerError)
}
