package route

import (
	"WebMsg/app/business/user"
	"github.com/gin-gonic/gin"
)

func AppInitRouters(r *gin.Engine) {

	r.POST("/admin-login", user.Login)
	r.POST("/admin-token", user.TestToken)
}
