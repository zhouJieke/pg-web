package user

import (
	"WebMsg/app/business/user/services"
	"WebMsg/route/middleware"
	"WebMsg/utils/results"
	"WebMsg/utils/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginStruct struct {
	UserName string `json:"userName" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

func Login(c *gin.Context) {
	var loginStruct LoginStruct
	_ = c.ShouldBindJSON(&loginStruct)
	msg, code := validator.Validate(loginStruct)
	if code != 200 {
		results.Error(c, msg, nil)
		return
	}
	userModel := services.QueryUserInfo(loginStruct.UserName, loginStruct.Password)
	if userModel.Id == 0 {
		results.Error(c, "用户名或者密码错误", nil)
		return
	}

	// 生成token存储redis，设置过期时间
	token := middleware.GenerateToken(&middleware.UserClaims{
		ID:             userModel.Id,
		StandardClaims: jwt.StandardClaims{},
	})

	results.Success(c, token)
}

func TestToken(c *gin.Context) {
	results.Success(c, nil)
}
