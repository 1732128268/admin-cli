package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func BaseRouter(group *gin.RouterGroup) {
	baseGroup := group.Group("base")
	userServer := system.UserServer{}
	{
		baseGroup.GET("captcha", system.Captcha)         // 验证码
		baseGroup.POST("login", userServer.Login)        // 登陆
		baseGroup.POST("/register", userServer.Register) // 注册
	}
}
