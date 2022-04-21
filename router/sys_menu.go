package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func MenuRouter(router *gin.RouterGroup) {
	menuRouter := router.Group("menu")
	menuServer := system.MenuService{}
	{
		menuRouter.GET("getBaseMenuList", menuServer.GetBaseMenuList)    // 分页获取基础menu列表
		menuRouter.POST("getBaseMenuById", menuServer.GetBaseMenuById)   // 根据id获取菜单
		menuRouter.POST("addBaseMenu", menuServer.AddBaseMenu)           // 新增菜单
		menuRouter.POST("deleteBaseMenu", menuServer.DeleteBaseMenu)     // 删除菜单
		menuRouter.POST("updateBaseMenu", menuServer.UpdateBaseMenu)     // 更新菜单
		menuRouter.POST("addMenuAuthority", menuServer.AddMenuAuthority) //	增加menu和角色关联关系
		//menuRouter.POST("getMenu", authorityMenuApi.GetMenu)                   // 获取菜单树
		//menuRouter.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   // 获取用户动态路由
		//menuRouter.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 获取指定角色menu

	}
}
