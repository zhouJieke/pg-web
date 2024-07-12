package route

import (
	"WebMsg/global"
	"WebMsg/route/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	// 设置日志级别
	gin.SetMode(global.App.Config.App.RunLogType)

	engine.Use(middleware.CustomRecovery())

	engine.Use(middleware.JwtVerify)
	//5.找不到路由
	engine.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.JSON(404, gin.H{"code": 404, "message": "您" + method + "请求地址：" + path + "不存在！"})
		return
	})

	return engine
}
