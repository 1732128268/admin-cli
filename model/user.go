package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        string          `json:"uuid" gorm:"comment:用户UUID"`                  // 用户UUID
	Username    string          `json:"userName" gorm:"comment:用户登录名"`               // 用户登录名
	Password    string          `json:"-"  gorm:"comment:用户登录密码"`                    // 用户登录密码
	NickName    string          `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`   // 用户昵称
	SideMode    string          `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"` // 用户侧边主题
	HeaderImg   string          `json:"headerImg" gorm:"comment:用户头像"`               // 用户头像 	// 活跃颜色
	Phone       string          `json:"phone"  gorm:"comment:用户手机号"`                 // 用户手机号
	Email       string          `json:"email"  gorm:"comment:用户邮箱"`                  // 用户邮箱
	Authorities []RoleAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}
