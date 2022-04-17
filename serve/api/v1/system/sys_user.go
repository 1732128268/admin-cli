package system

import (
	"admin-cli/global"
	"admin-cli/model"
	"admin-cli/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Login 用户登陆
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		global.ValidatorResponse(c, err)
		return
	}

	//查询用户
	if err := global.Db.Where("username = ?", user.Username).First(&user).Error; err != nil {
		logrus.Errorf("用户登陆失败，用户名不存在，%s", err)
		global.Response(c, nil, err)
		return
	}

	md5Password := utils.GenMd5(user.Password)
	if md5Password != user.Password {
		logrus.Error("用户登陆失败，密码错误")
		global.Response(c, nil, errors.New("用户名或密码错误"))
		return
	}

	//生成token
	token, err := utils.GenerateToken(user.ID, user.AuthorityId, user.Username)
	if err != nil {
		logrus.Errorf("generate token error:%v", err)
		global.Response(c, nil, err)
		return
	}

	global.Response(c, gin.H{
		"token": token,
		"user":  user,
	}, nil)

}

// Register 用户注册
func Register(c *gin.Context) {
	var (
		user model.User
	)
	//解析user数据
	if err := c.ShouldBind(&user); err != nil {
		logrus.Errorf("注册用户数据有误 err:%s", err)
		global.ValidatorResponse(c, err)
		return
	}
	//查询用户名是否存在
	err := global.Db.Where("username = ?", user.Username).First(&model.User{}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("查询用户名是否存在失败 err:%s", err)
		global.Response(c, nil, err)
		return
	}
	if err == nil {
		logrus.Errorf("用户名已存在 %v", user.Username)
		global.Response(c, nil, errors.New("用户名已存在"))
		return
	}

	var authorities []model.RoleAuthority
	for _, v := range user.AuthorityIds {
		authorities = append(authorities, model.RoleAuthority{
			AuthorityId: v,
		})
	}
	user.Password = utils.GenMd5(user.Password)
	user.Authorities = authorities
	if err := global.Db.Create(&user).Error; err != nil {
		logrus.Errorf("注册用户失败 err:%s", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"user": user,
	}, nil)

}
