package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func BaseRouter(group *gin.RouterGroup) {
	baseGroup := group.Group("base")
	{
		baseGroup.GET("captcha", system.Captcha) // 验证码
		baseGroup.POST("login", system.Login)    // 登陆
	}
}
