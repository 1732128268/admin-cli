package system

import (
	"admin-cli/global"
	"admin-cli/model"
	"admin-cli/model/gen"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthorityService struct{}

//查询二级权限
func findChildrenAuthority(authority *model.RoleAuthority) (err error) {
	err = global.Db.Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

// GetAuthorityList 分页获取角色列表
func (a *AuthorityService) GetAuthorityList(c *gin.Context) {
	size := c.DefaultQuery("size", "10")      //每页数
	current := c.DefaultQuery("current", "1") //当前页
	//总页数
	var total int64
	db := global.Db.Model(&model.RoleAuthority{})
	err := db.Where("parent_id = ?", "0").Count(&total).Error
	if err != nil {
		logrus.Errorf("获取角色列表数量失败，err:%v", err)
		global.Response(c, nil, err)
		return
	}
	var authority []model.RoleAuthority
	err = db.Scopes(gen.Paginate(cast.ToInt(current), cast.ToInt(size))).Where("parent_id = ?", "0").Find(&authority).Error
	if err != nil {
		logrus.Errorf("获取角色列表失败，err:%v", err)
		global.Response(c, nil, err)
		return
	}
	if len(authority) != 0 {
		//	处理二级权限
		for k := range authority {
			err = findChildrenAuthority(&authority[k])
			if err != nil {
				logrus.Errorf("获取角色二级列表失败，err:%v", err)
				global.Response(c, nil, err)
				return
			}
		}
	}

	global.Response(c, gin.H{
		"pageSize":    size, //每页数
		"authority":   authority,
		"pageTotal":   total,   //总页数
		"currentPage": current, //当前页
	}, nil)

}

// CreateAuthority 创建角色
func (a *AuthorityService) CreateAuthority(c *gin.Context) {
	var (
		authority model.RoleAuthority
		err       error
	)
	if err := c.ShouldBind(&authority); err != nil {
		global.ValidatorResponse(c, err)
		return
	}
	//检查角色id是否存在
	if err = global.Db.Where("authority_id = ?", authority.AuthorityId).First(&model.RoleAuthority{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("查询角色权限失败，%v", err)
		global.Response(c, nil, err)
		return
	}
	if err == nil {
		logrus.Errorf("角色权限ID已存在，%v", err)
		global.Response(c, nil, errors.New("角色权限ID已存在"))
		return
	}
	if err = global.Db.Create(&authority).Error; err != nil {
		logrus.Errorf("创建角色权限失败，%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"authority": authority,
	}, nil)
}

// DeleteAuthority 删除角色
func (a *AuthorityService) DeleteAuthority(c *gin.Context) {
	var authority model.RoleAuthority
	if err := c.ShouldBind(&authority); err != nil {
		logrus.Errorf("删除角色权限 ShouldBind err:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	//检查权限是否有用户使用
	if err := global.Db.Preload("Users").First(&authority).Error; err != nil {
		logrus.Errorf("查询角色权限失败，%v", err)
		global.Response(c, nil, err)
		return
	}

	if len(authority.Users) != 0 {
		logrus.Errorf("角色权限已被用户使用，%v", authority.Users)
		global.Response(c, nil, errors.New("角色权限已被用户使用"))
		return
	}
	if !errors.Is(global.Db.Where("authority_id = ?", authority.AuthorityId).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		logrus.Error("此角色权限已被用户使用")
		global.Response(c, nil, errors.New("此角色权限已被用户使用"))
		return
	}
	if !errors.Is(global.Db.Where("parent_id = ?", authority.AuthorityId).First(&model.RoleAuthority{}).Error, gorm.ErrRecordNotFound) {
		logrus.Errorf("此角色权限已被子角色使用，%v", authority.AuthorityId)
		global.Response(c, nil, errors.New("此角色权限已被子角色使用"))
		return
	}
	//删除角色权限
	db := global.Db.Preload("BaseMenus").Where("authority_id = ?", authority.AuthorityId).First(&authority)
	err := db.Unscoped().Delete(authority).Error
	if err != nil {
		logrus.Errorf("删除角色权限失败，%v", err)
		global.Response(c, nil, err)
		return
	}

	if len(authority.BaseMenus) > 0 {
		err = global.Db.Model(&authority).Association("BaseMenus").Delete(authority.BaseMenus)
		if err != nil {
			logrus.Errorf("删除角色权限菜单失败，%v", err)
			global.Response(c, nil, err)
			return
		}
	}
	CasbinServiceApp.ClearCasbin(0, authority.AuthorityId)
	global.Response(c, nil, nil)
}

// UpdateAuthority 更新角色信息
func (a *AuthorityService) UpdateAuthority(c *gin.Context) {
	var auth model.RoleAuthority
	if err := c.ShouldBind(&auth); err != nil {
		logrus.Errorf("更新角色信息 ShouldBind err:%v", err)
		global.Response(c, nil, err)
		return
	}
	err := global.Db.Where("authority_id = ?", auth.AuthorityId).First(&model.RoleAuthority{}).Updates(&auth).Error
	if err != nil {
		logrus.Errorf("更新角色信息 auth:%v err:%v", auth, err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"auth": auth,
	}, nil)

}

// GetAuthority 根据id获取当前角色的数据
func (a *AuthorityService) GetAuthority(c *gin.Context) {
	authorityId := c.Query("authority_id")
	var auth model.RoleAuthority
	err := global.Db.Preload(clause.Associations).Where("authority_id = ?", authorityId).First(&auth).Error
	if err != nil {
		logrus.Errorf("GetAuthority 根据authority_id:%v err:%v", authorityId, err)
		global.Response(c, nil, err)
		return
	}
	//获取二级权限
	err = findChildrenAuthority(&auth)
	if err != nil {
		logrus.Errorf("GetAuthority 获取角色二级列表失败，err:%v", err)
		global.Response(c, nil, err)
		return
	}

	global.Response(c, gin.H{
		"auth": auth,
	}, nil)
}
