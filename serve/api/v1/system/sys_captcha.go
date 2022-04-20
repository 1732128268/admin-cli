package system

import (
	"admin-cli/global"
	"admin-cli/utils/captcha"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
var store = captcha.NewDefaultRedisStore()

//var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c)) // v8下使用redis
	//cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		logrus.Errorf("获取验证码失败，err: %v", err)
		global.Response(c, nil, errors.New("获取验证码失败"))
	} else {
		global.Response(c, gin.H{
			"id":     id,
			"img":    b64s,
			"length": global.Config.Captcha.KeyLong,
		}, nil)
	}
}
