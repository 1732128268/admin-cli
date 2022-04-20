package initialize

import (
	"admin-cli/router"
	"github.com/gin-gonic/gin"
)

// Router 注册路由
func Router(r *gin.Engine) {
	group := r.Group("")
	router.BaseRouter(group)      // 基础路由
	router.UserRouter(group)      // 用户
	router.AuthorityRouter(group) // 角色
	router.MenuRouter(group)      //权限菜单
	router.CasbinRouter(group)    //casbin获取

}
