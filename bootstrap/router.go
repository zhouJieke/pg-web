package bootstrap

import (
	"WebMsg/global"
	"WebMsg/route"
)

// RunServer 启动服务器。
func RunServer() {
	r := route.InitRouter()
	// 注册路由
	route.AppInitRouters(r)
	_ = r.Run(":" + global.App.Config.App.Port)
}
