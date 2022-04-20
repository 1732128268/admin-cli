package model

type SysUseAuthority struct {
	UserId uint `gorm:"column:user_id"`
	//RoleAuthorityId          string `gorm:"column:role_authority_id"`
	RoleAuthorityAuthorityId string `gorm:"column:role_authority_authority_id"`
}

func (s *SysUseAuthority) TableName() string {
	return "sys_user_authority"
}

type SysAuthorityMenus struct {
	SysBaseMenuId uint `gorm:"column:sys_base_menu_id"`
	//RoleAuthorityId          string `gorm:"column:role_authority_id"`
	RoleAuthorityAuthorityId string `gorm:"column:role_authority_authority_id"`
}

func (s *SysAuthorityMenus) TableName() string {
	return "sys_authority_menus"
}
