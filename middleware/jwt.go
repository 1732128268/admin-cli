package middleware

import (
	"admin-cli/global"
	"admin-cli/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetClaims(c *gin.Context) (*utils.AuthClaims, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("请携带token")
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("请求头中auth格式有误")
	}
	jwt, err := utils.ParseToken(parts[1])
	return jwt, err
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.DebugMode == global.Config.HttpConfig.Mode {
			c.Next()
			return
		}
		jwt, err := GetClaims(c)
		if err != nil {
			global.Response(c, "", err)
			c.Abort()
			return
		}
		c.Set("id", jwt.Id)
		c.Set("name", jwt.Name)
		c.Set("roleId", jwt.RoleId)

	}
}
