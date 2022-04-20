package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func MenuRouter(router *gin.RouterGroup) {
	menuRouter := router.Group("menu")
	{
		menuRouter.GET("getBaseMenuList", system.GetBaseMenuList)    // 分页获取基础menu列表
		menuRouter.POST("getBaseMenuById", system.GetBaseMenuById)   // 根据id获取菜单
		menuRouter.POST("addBaseMenu", system.AddBaseMenu)           // 新增菜单
		menuRouter.POST("deleteBaseMenu", system.DeleteBaseMenu)     // 删除菜单
		menuRouter.POST("updateBaseMenu", system.UpdateBaseMenu)     // 更新菜单
		menuRouter.POST("addMenuAuthority", system.AddMenuAuthority) //	增加menu和角色关联关系
		//menuRouter.POST("getMenu", authorityMenuApi.GetMenu)                   // 获取菜单树
		//menuRouter.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   // 获取用户动态路由
		//menuRouter.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 获取指定角色menu

	}
}
