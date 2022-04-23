package middleware

import (
	"admin-cli/global"
	"errors"
	"github.com/gin-gonic/gin"
)

//令牌捅限流

func BucketRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !global.Config.HttpConfig.OpenRatelimit {
			c.Next()
		}
		if global.Bucket.TakeAvailable(1) < 1 {
			c.Abort()
			global.Response(c, nil, errors.New("服务器繁忙，请稍后再试"))
			return
		}
		c.Next()

	}
}
