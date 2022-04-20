package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup) {
	UserRouter := group.Group("user")
	//UserRouter.Use(middleware.Auth())
	{
		UserRouter.GET("/userList", system.UserList)
		UserRouter.POST("/register", system.Register) // 注册
	}
}
