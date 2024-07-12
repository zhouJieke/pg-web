package middleware

import (
	"WebMsg/resource/config"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type UserClaims struct {
	ID       int    `json:"id"`
	Openid   string `json:"openid"` //微信openid
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f489")
	// effectTime = 2 * time.Minute //两分钟
)

func GenerateToken(claims *UserClaims) string {
	claims.IssuedAt = time.Now().Unix()
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		panic(err)
	}
	_, _ = config.GetRedisHelper().Set(context.Background(), sign, claims.ID, 30*time.Minute).Result()
	return sign
}

// JwtVerify 验证JWT（JSON Web Token）的有效性。
func JwtVerify(c *gin.Context) {
	var noVerifyTokenArr = []string{
		"/admin-login",
	}
	if IsContain(noVerifyTokenArr, c.Request.URL.Path) {
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		panic("token 不存在")
	}

	// 校验token是否过期
	if _, err := config.GetRedisHelper().Get(context.Background(), token).Result(); err != nil {
		panic("token已过期")
	}

	//验证token，并存储在请求中
	c.Set("user", ParseToken(token))
}

// ParseToken 解析token（JSON Web Token）的有效性。

func ParseToken(tokenString string) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic("token解析错误")
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("The token is invalid")
	}

	// 续约token
	_, _ = config.GetRedisHelper().Set(context.Background(), tokenString, claims.ID, 30*time.Minute).Result()

	return claims
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
