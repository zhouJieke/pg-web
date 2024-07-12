package main

import (
	"WebMsg/bootstrap"
	"WebMsg/global"
	"WebMsg/resource/config"
	"context"
	"go.uber.org/zap"
)

var (
	ctx = context.Background()
)

func main() {
	// 初始化配置
	global.App.Config.InitializeConfig()
	// 初始化日志配置
	global.App.Log = bootstrap.InitializeLog()

	// 初始化mysql
	config.InitMysql(global.App.Config.MysqlConfig)

	// 初始化redis
	rdb := config.NewRedisHelper(global.App.Config.RedisConfig)

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		global.App.Log.Error("redis连接失败", zap.Error(err))
		return
	}
	// 启动服务器
	bootstrap.RunServer()
}
