package game

import (
	"WebMsg/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// Response 请求失败返回
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ServerError(c *gin.Context, err interface{}) {
	conf := global.App.Config
	msg := "内部服务器错误"
	if os.Getenv(gin.EnvGinMode) != gin.ReleaseMode && reflect.TypeOf(err).Name() == "string" {
		msg = err.(string)
	} else {
		if conf.App.Env != "pro" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
			if _, ok := err.(error); ok {
				msg = err.(error).Error()
			}
		} else {
			str := fmt.Sprintf("内部服务器错误： %s\n", err.(error).Error()) //拼接字符串
			global.App.Log.Error(str)
		}
	}
	//判断错误类型
	if res := strings.Contains(msg, "token is expired by"); res { //token超时
		c.JSON(200, Response{
			401,
			nil,
			msg,
		})
	} else if res := strings.Contains(msg, "invalid memory address or nil pointer dereference"); res { //数据库链接失败
		//model.MyInit(3) //重连数据库-初始化数据
		c.JSON(http.StatusInternalServerError, Response{1,
			"可能是数据库链接失败，请查看数据库链接是否正常",
			msg + "，可能是数据库链接失败，请查看数据库配置及是否启动，再刷新试试！",
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{1,
			nil,
			msg,
		})
	}
	c.Abort()
}
