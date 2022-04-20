package model

import (
	"gorm.io/gorm"
	"time"
)

// RoleAuthority 角色表
type RoleAuthority struct {
	CreatedAt   time.Time       // 创建时间
	UpdatedAt   time.Time       // 更新时间
	DeletedAt   *time.Time      `sql:"index"`
	AuthorityId string          `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	Name        string          `gorm:"type:varchar(100);" json:"name"`                                      //角色名称
	ParentId    string          `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	Children    []RoleAuthority `json:"children" gorm:"-"`
	BaseMenus   []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users       []User          `json:"-" gorm:"many2many:sys_user_authority;"`
}

// SysBaseMenu 权限菜单
type SysBaseMenu struct {
	gorm.Model
	ParentId    string          `json:"parentId" gorm:"comment:父菜单ID"` // 父菜单ID
	Title       string          `json:"title" gorm:"comment:菜单名"`      // 菜单名
	Icon        string          `json:"icon" gorm:"comment:菜单图标"`      // 菜单图标
	Path        string          `json:"path" gorm:"comment:路由path"`    // 路由path
	Name        string          `json:"name" gorm:"comment:路由name"`    // 路由name`
	Method      string          `json:"method" gorm:"comment:请求方法"`    //请求方法
	Sort        int             `json:"sort" gorm:"comment:排序"`        //排序
	Children    []SysBaseMenu   `json:"children" gorm:"-"`
	Authorities []RoleAuthority `json:"authorities" gorm:"many2many:sys_authority_menus;"`
}
