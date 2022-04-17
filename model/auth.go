package model

import "gorm.io/gorm"

type RoleAuthority struct {
	gorm.Model
	AuthorityId string          `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	Name        string          `gorm:"type:varchar(100);"`                                                  //角色名称
	ParentId    string          `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	Children    []RoleAuthority `json:"children" gorm:"-"`
	BaseMenus   []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users       []User          `json:"-" gorm:"many2many:sys_user_authority;"`
}

// SysBaseMenu 权限菜单
type SysBaseMenu struct {
	gorm.Model
	ParentId    string          `json:"parentId" gorm:"comment:父菜单ID"` // 父菜单ID
	Path        string          `json:"path" gorm:"comment:路由path"`    // 路由path
	Name        string          `json:"name" gorm:"comment:路由name"`    // 路由name`
	Method      string          `json:"method" gorm:"comment:请求方法"`    //请求方法
	Sort        int             `json:"sort" gorm:"comment:排序"`        //排序
	Authorities []RoleAuthority `json:"authorities" gorm:"many2many:sys_authority_menus;"`
}
