package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func AuthorityRouter(router *gin.RouterGroup) {
	authority := router.Group("authority")
	{
		authority.GET("getAuthorityList", system.GetAuthorityList) // 获取角色列表
		authority.GET("getAuthority", system.GetAuthority)         // 根据角色id获取角色信息
		authority.POST("createAuthority", system.CreateAuthority)  // 创建角色
		authority.POST("deleteAuthority", system.DeleteAuthority)  // 删除角色
		authority.POST("updateAuthority", system.UpdateAuthority)  // 更新角色
	}

}
