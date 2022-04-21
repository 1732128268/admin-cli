package initialize

import (
	"admin-cli/global"
	"admin-cli/model"
	"admin-cli/utils"
	"gorm.io/gorm"
)

// Admin 初始化用户 添加admin
func Admin() error {
	var user model.User
	err := global.Db.Where("username = ?", "admin").
		Where("authority_id = ?", 888).
		Where("nick_name = ?", "超级管理员").
		First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == nil {
		return nil
	}
	user = model.User{
		UUID:     utils.GetUid(),
		Username: "admin",
		Password: "d41d8cd98f00b204e9800998ecf8427e",
		NickName: "超级管理员",
		Phone:    "17611111111",
	}
	if err := global.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
