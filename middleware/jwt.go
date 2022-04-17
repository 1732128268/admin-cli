package middleware

import (
	"admin-cli/global"
	"admin-cli/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			global.Response(c, "", errors.New("请携带token"))
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			global.Response(c, "", errors.New("请求头中auth格式有误"))
			c.Abort()
			return
		}
		jwt, err := utils.ParseToken(parts[1])
		if err != nil {
			global.Response(c, "", err)
			c.Abort()
			return
		}
		c.Set("id", jwt.Id)
		c.Set("name", jwt.Name)
		c.Set("roleId", jwt.RoleId)
		c.Next()
	}
}
