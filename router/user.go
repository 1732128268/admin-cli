package router

import (
	"admin-cli/middleware"
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup) {
	UserRouter := group.Group("user")
	UserRouter.Use(middleware.Auth())
	{
		UserRouter.POST("/register", system.Register) // 注册
	}
}
