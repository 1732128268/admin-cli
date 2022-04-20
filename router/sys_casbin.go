package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func CasbinRouter(router *gin.RouterGroup) {
	casbinRouter := router.Group("casbin")
	{
		casbinRouter.POST("getPolicyPathByAuthorityId", system.GetPolicyPathByAuthorityId)
	}
}
