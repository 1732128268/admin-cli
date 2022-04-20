package serve

import (
	"admin-cli/config"
	"admin-cli/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/threading"
	"net/http"
)

// StartHttp gin优雅重启
func StartHttp() (*http.Server, *gin.Engine) {
	//设置模式
	gin.SetMode(global.Config.HttpConfig.Mode)
	//初始化gin
	router := gin.Default()
	conf := config.GetConfig()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.HttpConfig.Port),
		Handler: router,
	}
	threading.GoSafe(func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	})
	fmt.Println(fmt.Sprintf("启动成功，监听端口：%d", conf.HttpConfig.Port))
	return srv, router
}
