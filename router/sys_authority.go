package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func AuthorityRouter(router *gin.RouterGroup) {
	authority := router.Group("authority")
	authorityService := system.AuthorityService{}
	{
		authority.GET("getAuthorityList", authorityService.GetAuthorityList) // 获取角色列表
		authority.GET("getAuthority", authorityService.GetAuthority)         // 根据角色id获取角色信息
		authority.POST("createAuthority", authorityService.CreateAuthority)  // 创建角色
		authority.POST("deleteAuthority", authorityService.DeleteAuthority)  // 删除角色
		authority.POST("updateAuthority", authorityService.UpdateAuthority)  // 更新角色
	}

}
