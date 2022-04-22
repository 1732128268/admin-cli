package middleware

import (
	"admin-cli/global"
	"admin-cli/serve/api/v1/system"
	"errors"
	"github.com/gin-gonic/gin"
)

var casbinService = system.CasbinService{}

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.DebugMode == global.Config.HttpConfig.Mode {
			c.Next()
			return
		}
		waitUse, _ := GetClaims(c)

		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		e := casbinService.Casbin()
		// 判断策略中是否存在
		var success bool
		for _, v := range waitUse.RoleId {
			success, _ = e.Enforce(v, obj, act)
			if success {
				break
			}
		}
		if success {
			c.Next()
		} else {
			global.Response(c, nil, errors.New("权限不足"))
			c.Abort()
			return
		}
	}
}
