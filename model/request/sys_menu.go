package request

import "admin-cli/model"

// AddMenuAuthorityInfo 角色权限关联
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu `json:"menus"`
	AuthorityId string              `json:"authorityId"` // 角色ID
}
