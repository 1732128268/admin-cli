package router

import (
	"admin-cli/middleware"
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func SystemRouter(router *gin.RouterGroup) {
	sysRouter := router.Group("system").Use(middleware.OperationRecord())
	systemApi := system.SystemApi{}
	{
		sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("getServerInfo", systemApi.GetServerInfo)     // 获取服务器信息
	}
}
