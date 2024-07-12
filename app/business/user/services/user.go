package services

import (
	"WebMsg/model"
	"WebMsg/resource/config"
)

// QueryUserInfo
// 通过账户和密码查询用户是否过期
func QueryUserInfo(userName string, password string) model.AdminUsers {
	var userModel model.AdminUsers
	config.Engin.Where("user_name = ? AND password = ?", userName, password).First(&userModel)
	return userModel
}
