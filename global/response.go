package global

import (
	"admin-cli/vaildator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"net/http"
)

var (
	Db    *gorm.DB
	Redis *redis.Client
)

const (
	LIKE     = "%%%v%%"
	CronSpec = "%v %v %v %v *"
)

func Response(c *gin.Context, msg interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"err":  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "0",
			"msg":  msg,
		})
		return
	}

}

func ValidatorResponse(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  err.Error(),
		})
		return

	}
	// validator.ValidationErrors类型错误则进行翻译
	c.JSON(http.StatusOK, gin.H{
		"code": "1",
		"msg":  vaildator.RemoveTopStruct(errs.Translate(vaildator.Trans)),
	})
}
