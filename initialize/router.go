package initialize

import (
	"admin-cli/router"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	group := r.Group("")

	router.UserRouter(group)
}
