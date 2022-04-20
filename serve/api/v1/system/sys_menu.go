package system

import (
	"admin-cli/global"
	"admin-cli/model"
	"admin-cli/model/gen"
	"admin-cli/model/request"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

//二级权限菜单
func findChildrenBaseMenu(baseMenu *model.SysBaseMenu) (err error) {
	err = global.Db.Where("parent_id = ?", baseMenu.ID).Find(&baseMenu.Children).Error
	if len(baseMenu.Children) > 0 {
		for k := range baseMenu.Children {
			err = findChildrenBaseMenu(&baseMenu.Children[k])
		}
	}
	return err
}

// GetBaseMenuList 获取所有权限菜单
func GetBaseMenuList(c *gin.Context) {
	size := c.DefaultQuery("size", "10")      //每页数
	current := c.DefaultQuery("current", "1") //当前页
	//总页数
	var total int64
	//获取一级权限数量
	global.Db.Model(&model.SysBaseMenu{}).Where("parent_id = ?", 0).Count(&total)
	//获取一级权限列表
	var baseMenuList []model.SysBaseMenu
	global.Db.Scopes(gen.Paginate(cast.ToInt(current), cast.ToInt(size))).Where("parent_id = ?", 0).Find(&baseMenuList)
	if len(baseMenuList) > 0 {
		//处理二级菜单
		for k := range baseMenuList {
			err := findChildrenBaseMenu(&baseMenuList[k])
			if err != nil {
				logrus.Errorf("findChildrenBaseMenu error: %v", err)
				global.Response(c, nil, err)
				return
			}
		}
	}
	global.Response(c, gin.H{
		"total":   total,   //总条数
		"current": current, //当前页
		"size":    size,    //每页数
		"list":    baseMenuList,
	}, nil)
}

// GetBaseMenuById 根据id获取菜单
func GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	if err := c.ShouldBind(&idInfo); err != nil {

	}
	//根据id查询权限菜单
	var baseMenu model.SysBaseMenu
	err := global.Db.Where("id = ?", idInfo.ID).First(&baseMenu).Error
	if err != nil {
		logrus.Errorf("GetBaseMenuById id:%v err:%v", idInfo.ID, err)
		global.Response(c, nil, err)
		return
	}

	//获取2级权限菜单
	err = findChildrenBaseMenu(&baseMenu)
	if err != nil {
		logrus.Errorf("GetBaseMenuById findChil	drenBaseMenu err:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"menu": baseMenu,
	}, nil)
}

// AddBaseMenu 新增权限菜单
func AddBaseMenu(c *gin.Context) {
	var (
		menu model.SysBaseMenu
		err  error
	)
	if err := c.ShouldBind(&menu); err != nil {
		logrus.Errorf("AddBaseMenu Bind Error:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	//检验菜单名称是否重复
	if err = global.Db.Where("name = ?", menu.Name).First(&model.SysBaseMenu{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("AddBaseMenu Check Name Error:%v", err)
		global.Response(c, nil, err)
	}

	if err == nil {
		logrus.Errorf("AddBaseMenu Check Name:%v 已存在", menu.Name)
		global.Response(c, nil, errors.New("菜单名称已存在"))
		return
	}

	if err = global.Db.Create(&menu).Error; err != nil {
		logrus.Errorf("AddBaseMenu Create Error:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"menu": menu,
	}, nil)
	return

}

// DeleteBaseMenu 删除权限菜单
func DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	if err := c.ShouldBind(&menu); err != nil {
		logrus.Errorf("DeleteBaseMenu ShouldBind err:%v", err)
		global.Response(c, nil, err)
		return
	}
	//查询是否有子菜单
	var baseMenu model.SysBaseMenu
	err := global.Db.Where("parent_id = ?", menu.ID).First(&baseMenu).Error
	if err == nil {
		logrus.Error("DeleteBaseMenu存在子权限")
		global.Response(c, nil, errors.New("存在子权限"))
		return
	}
	err = global.Db.Preload("Authorities").Where("id = ?", menu.ID).First(&baseMenu).Delete(&baseMenu).Error
	if err != nil {
		logrus.Errorf("DeleteBaseMenu Delete Error:%v", err)
		global.Response(c, nil, err)
		return
	}
	if len(baseMenu.Authorities) > 0 {
		err = global.Db.Model(&baseMenu).Association("Authorities").Delete(&baseMenu.Authorities)
		if err != nil {
			logrus.Errorf("DeleteBaseMenu Delete Authorities Error:%v", err)
			global.Response(c, nil, err)
			return
		}
	}
	CasbinServiceApp.ClearCasbin(1, baseMenu.Path, baseMenu.Method)
	global.Response(c, nil, nil)

}

// UpdateBaseMenu 更新权限菜单
func UpdateBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	if err := c.ShouldBind(&menu); err != nil {
		logrus.Errorf("UpdateBaseMenu ShouldBind err:%v", err)
		global.Response(c, nil, err)
		return
	}
	var baseMenu model.SysBaseMenu
	err := global.Db.Where("id = ?", menu.ID).First(&baseMenu).Error
	if err != nil {
		logrus.Errorf("UpdateBaseMenu First id:%v err:%v", menu.ID, err)
		global.Response(c, nil, err)
		return
	}
	//查询更新的权限名称是否重复
	if err = global.Db.Where("name = ? ", menu.Name).First(&model.SysBaseMenu{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("UpdateBaseMenu Check Name Error:%v", err)
		global.Response(c, nil, err)
	}
	if err == nil {
		logrus.Errorf("更新的名称已存在 name:%v", menu.Name)
		global.Response(c, nil, errors.New("更新的名称已存在"))
		return
	}

	if baseMenu.Path != menu.Path || baseMenu.Method != menu.Method {
		if !errors.Is(global.Db.Where("path = ? AND method = ?", menu.Path, menu.Method).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
			logrus.Errorf("更新的路径和方法已存在 path:%v method:%v", menu.Path, menu.Method)
			global.Response(c, nil, errors.New("更新的路径和方法已存在"))
			return
		}
	}
	err = CasbinServiceApp.UpdateCasbinApi(baseMenu.Path, menu.Path, baseMenu.Method, menu.Method)
	if err != nil {
		logrus.Errorf("UpdateBaseMenu UpdateCasbinApi Error:%v", err)
		global.Response(c, nil, err)
		return
	}
	baseMenu.Name = menu.Name
	baseMenu.Path = menu.Path
	baseMenu.Icon = menu.Icon
	baseMenu.Sort = menu.Sort
	baseMenu.ParentId = menu.ParentId
	baseMenu.Method = menu.Method
	baseMenu.Title = menu.Title
	err = global.Db.Model(&baseMenu).Updates(&baseMenu).Error
	if err != nil {
		logrus.Errorf("UpdateBaseMenu Update Error:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"menu": baseMenu,
	}, nil)
}

// AddMenuAuthority 增加menu和角色关联关系
func AddMenuAuthority(c *gin.Context) {
	var authorityMenu request.AddMenuAuthorityInfo
	if err := c.ShouldBind(&authorityMenu); err != nil {
		logrus.Errorf("AddMenuAuthority ShouldBind err:%v", err)
		global.Response(c, nil, err)
		return
	}
	var authority model.RoleAuthority
	err := global.Db.Preload("BaseMenus").First(&authority, "authority_id = ?", authorityMenu.AuthorityId).Error
	if err != nil {
		logrus.Errorf("AddMenuAuthority First authority_id:%v err:%v", authorityMenu.AuthorityId, err)
		global.Response(c, nil, err)
		return
	}
	//开启事务
	tx := global.Db.Begin()
	if tx.Error != nil {
		logrus.Errorf("AddMenuAuthority Begin Error:%v", tx.Error)
		global.Response(c, nil, tx.Error)
		return
	}
	// 更新casbin
	var casbinInfos []request.CasbinInfo
	for index, _ := range authorityMenu.Menus {
		menu := authorityMenu.Menus[index]
		casbinInfos = append(casbinInfos, request.CasbinInfo{
			Path:   menu.Path,
			Method: menu.Method,
		})
	}
	err = CasbinServiceApp.UpdateCasbin(authority.AuthorityId, casbinInfos)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("AddMenuAuthority UpdateCasbin Error:%v", err)
		global.Response(c, nil, err)
		return
	}
	err = global.Db.Model(&authority).Association("BaseMenus").Replace(&authorityMenu.Menus)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("AddMenuAuthority Replace err:%v", err)
		global.Response(c, nil, err)
		return
	}

	if err = tx.Commit().Error; err != nil {
		logrus.Errorf("AddMenuAuthority Commit err:%v", err)
		global.Response(c, nil, err)
		return
	}

	global.Response(c, gin.H{
		"authority": authority,
	}, nil)
}
