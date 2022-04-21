package router

import (
	"admin-cli/middleware"
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup) {
	userRouter := group.Group("user")
	userRouter.Use(middleware.Auth())
	userServer := system.UserServer{}
	{
		userRouter.GET("userList", userServer.UserList)                      // 获取用户列表
		userRouter.GET("getUserInfo", userServer.GetUserInfo)                // 更具id获取用户信息
		userRouter.POST("changePassword", userServer.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthorities", userServer.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("deleteUser", userServer.DeleteUser)                 // 删除用户
		userRouter.POST("setUserInfo", userServer.SetUserInfo)               // 设置用户信息
		userRouter.POST("resetPassword", userServer.ResetPassword)           // 重置密码
	}
}
