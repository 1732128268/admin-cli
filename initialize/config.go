package initialize

import (
	"admin-cli/config"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitConfig 初始化config
func InitConfig() error {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("配置文件读取err", err)
		return err
	}
	var configData config.Config
	if err = viper.Unmarshal(&configData); err != nil {
		logrus.Error("配置文件 Unmarshal err", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Info("配置文件发生更新")
		if err := viper.Unmarshal(&configData); err != nil {
			logrus.Infof("配置文件更新解析失败,err:%v", err)
		}
	})
	config.SetConfig(&configData)
	return nil
}
