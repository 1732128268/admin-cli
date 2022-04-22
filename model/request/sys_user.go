package request

import "admin-cli/model"

type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Register 注册
type Register struct {
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	NickName     string   `json:"nickName"`
	HeaderImg    string   `json:"headerImg"`
	Phone        string   `json:"phone"` // 用户手机号
	Email        string   `json:"email"`
	AuthorityIds []string `json:"authorityIds"`
}

// ChangePasswordStruct 修改密码
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// SetUserAuth 角色id
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}

// SetUserAuthorities 用户设置角色id
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `json:"id"`           // 主键ID
	NickName     string                `json:"nickName"`     // 用户昵称
	Phone        string                `json:"phone"`        // 用户角色ID
	AuthorityIds []string              `json:"authorityIds"` // 角色ID
	Email        string                `json:"email"  `      // 用户邮箱
	HeaderImg    string                `json:"headerImg"`    // 用户头像
	Authorities  []model.RoleAuthority `json:"-"`
}
