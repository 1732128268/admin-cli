package system

import (
	"admin-cli/global"
	"admin-cli/model"
	"admin-cli/model/gen"
	"admin-cli/model/request"
	"admin-cli/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type UserServer struct{}

// Login 用户登陆
func (u *UserServer) Login(c *gin.Context) {
	var login request.Login
	if err := c.ShouldBind(&login); err != nil {
		global.ValidatorResponse(c, err)
		return
	}
	//开启验证码校验
	if global.Config.HttpConfig.OpenCaptcha {
		if global.Config.HttpConfig.OpenRedis {
			if !store.UseWithCtx(c).Verify(login.CaptchaId, login.Captcha, true) {
				logrus.Errorf("验证码错误")
				global.Response(c, nil, errors.New("验证码错误"))
				return
			}
		} else {
			if !store.Verify(login.CaptchaId, login.Captcha, true) {
				logrus.Errorf("验证码错误")
				global.Response(c, nil, errors.New("验证码错误"))
				return
			}
		}
	}

	//查询用户
	var user model.User
	if err := global.Db.Where("username = ?", login.Username).Preload("Authorities").First(&user).Error; err != nil {
		logrus.Errorf("用户登陆失败，用户名不存在，%s", err)
		global.Response(c, nil, err)
		return
	}

	md5Password := utils.GenMd5(login.Password)
	if md5Password != user.Password {
		logrus.Error("用户登陆失败，密码错误")
		global.Response(c, nil, errors.New("用户名或密码错误"))
		return
	}
	//权限id
	var roleIds []string
	for _, role := range user.Authorities {
		roleIds = append(roleIds, role.AuthorityId)
	}
	//生成token
	token, err := utils.GenerateToken(user.ID, user.Username, roleIds)
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

// UserList 用户列表
func (u *UserServer) UserList(c *gin.Context) {
	size := c.DefaultQuery("size", "10")      //每页数
	current := c.DefaultQuery("current", "1") //当前页
	//总页数
	var total int64
	//获取一级权限数量
	global.Db.Model(&model.User{}).Count(&total)
	//用户列表
	var users []model.User
	//查询用户列表
	if err := global.Db.Scopes(gen.Paginate(cast.ToInt(current), cast.ToInt(size))).Preload("Authorities").Find(&users).Error; err != nil {
		logrus.Errorf("查询用户列表失败，%s", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"total":   total,   //总条数
		"current": current, //当前页
		"size":    size,    //每页数
		"users":   users,
	}, nil)
}

// Register 用户注册
func (u *UserServer) Register(c *gin.Context) {
	var (
		register request.Register
	)
	//解析user数据
	if err := c.ShouldBind(&register); err != nil {
		logrus.Errorf("注册用户数据有误 err:%s", err)
		global.ValidatorResponse(c, err)
		return
	}
	//查询用户名是否存在
	err := global.Db.Where("username = ?", register.Username).First(&model.User{}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("查询用户名是否存在失败 err:%s", err)
		global.Response(c, nil, err)
		return
	}
	if err == nil {
		logrus.Errorf("用户名已存在 %v", register.Username)
		global.Response(c, nil, errors.New("用户名已存在"))
		return
	}

	var authorities []model.RoleAuthority
	for _, v := range register.AuthorityIds {
		authorities = append(authorities, model.RoleAuthority{
			AuthorityId: v,
		})
	}
	var user model.User
	user.UUID = utils.GetUid()
	user.Username = register.Username
	user.NickName = register.NickName
	user.Password = utils.GenMd5(register.Password)
	user.HeaderImg = register.HeaderImg
	user.Phone = register.Phone
	user.Email = register.Email
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

// ChangePassword 用户修改密码
func (u *UserServer) ChangePassword(c *gin.Context) {
	var user request.ChangePasswordStruct
	if err := c.ShouldBind(&user); err != nil {
		logrus.Errorf("ChangePassword ShouldBind data:%v err:%v", user, err)
		global.ValidatorResponse(c, err)
		return
	}

	var userInfo model.User
	oldPwd := utils.GenMd5(user.Password)
	newPwd := utils.GenMd5(user.NewPassword)
	if err := global.Db.Where("username = ?", user.Username).Where("password = ?", oldPwd).First(&userInfo).Update("password", newPwd).Error; err != nil {
		logrus.Errorf("ChangePassword 修改密码 err:%v", err)
		global.Response(c, nil, errors.New("修改失败，原密码与当前账户不符"))
		return
	}
	global.Response(c, nil, nil)
}

//
func setUserAuthorities(id uint, authorityIds []string) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&[]model.SysUseAuthority{}, "user_id = ?", id).Error
		if err != nil {
			return err
		}
		var useAuthority []model.SysUseAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, model.SysUseAuthority{
				UserId:                   id,
				RoleAuthorityAuthorityId: v,
			})
		}
		tx = tx.Create(&useAuthority)
		if tx.Error != nil {
			return tx.Error
		}
		// 返回 nil 提交事务
		return nil
	})
}

// SetUserAuthorities 设置用户权限
func (u *UserServer) SetUserAuthorities(c *gin.Context) {
	var sua request.SetUserAuthorities
	if err := c.ShouldBind(&sua); err != nil {
		logrus.Errorf("SetUserAuthorities ShouldBind data:%v err:%v", sua, err)
		global.ValidatorResponse(c, err)
		return
	}
	err := setUserAuthorities(sua.ID, sua.AuthorityIds)
	if err != nil {
		logrus.Errorf("SetUserAuthorities 修改权限 err:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, nil, nil)
}

// DeleteUser 删除用户
func (u *UserServer) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	if err := c.ShouldBind(&reqId); err != nil {
		logrus.Errorf("DeleteUser ShouldBind err:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	var user model.User
	err := global.Db.Where("id = ?", reqId.ID).Delete(&user).Error
	if err != nil {
		logrus.Errorf("DeleteUser 删除用户 err:%v", err)
		global.Response(c, nil, err)
		return
	}
	err = global.Db.Delete(&[]model.SysUseAuthority{}, "user_id = ?", reqId.ID).Error
	if err != nil {
		logrus.Errorf("DeleteUser 删除用户权限 err:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, nil, nil)
}

// SetUserInfo 更新用户信息
func (u *UserServer) SetUserInfo(c *gin.Context) {
	var info request.ChangeUserInfo
	if err := c.ShouldBind(&info); err != nil {
		logrus.Errorf("SetUserInfo ShouldBind data:%v err:%v", info, err)
		global.ValidatorResponse(c, err)
		return
	}

	if len(info.AuthorityIds) != 0 {
		err := setUserAuthorities(info.ID, info.AuthorityIds)
		if err != nil {
			logrus.Errorf("SetUserInfo 修改权限 userId:%v authIds:%v err:%v", info.ID, info.AuthorityIds, err)
			global.Response(c, nil, err)
			return
		}
	}

	user := model.User{
		Model:     gorm.Model{ID: info.ID},
		NickName:  info.NickName,
		HeaderImg: info.HeaderImg,
		Phone:     info.Phone,
		Email:     info.Email,
	}

	err := global.Db.Updates(&user).Error
	if err != nil {
		logrus.Errorf("SetUserInfo 更新用户信息 err:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"user": user,
	}, nil)

}

// GetUserInfo 获取用户信息
func (u *UserServer) GetUserInfo(c *gin.Context) {
	var reqId request.GetById
	if err := c.ShouldBind(&reqId); err != nil {
		logrus.Errorf("GetUserInfo ShouldBind err:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	var user model.User
	err := global.Db.Preload("Authorities").Where("id = ?", reqId.ID).First(&user).Error
	if err != nil {
		logrus.Errorf("GetUserInfo 查询用户信息 err:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"user": user,
	}, nil)

}

// ResetPassword 重置用户密码
func (u *UserServer) ResetPassword(c *gin.Context) {
	var reqId request.GetById
	if err := c.ShouldBind(&reqId); err != nil {
		logrus.Errorf("GetUserInfo ShouldBind err:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	var user model.User
	err := global.Db.Where("id = ?", reqId.ID).First(&user).Update("password", utils.GenMd5("123456")).Error
	if err != nil {
		logrus.Errorf("ResetPassword 重置用户密码 err:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, nil, nil)
}
