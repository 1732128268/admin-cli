package system

import (
	config2 "admin-cli/config"
	"admin-cli/global"
	"admin-cli/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type SystemApi struct{}

// GetSystemConfig 获取配置文件内容
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	global.Response(c, gin.H{
		"config": global.Config,
	}, nil)
}

// SetSystemConfig 设置配置文件内容
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var config config2.Config
	if err := c.ShouldBind(&config); err != nil {
		logrus.Errorf("设置配置文件 ShouldBind err:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	cs := utils.StructToMap(config)
	for k, v := range cs {
		viper.Set(k, v)
	}
	err := viper.WriteConfig()
	if err != nil {
		logrus.Errorf("设置配置文件 WriteConfig err:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, nil, nil)
}

// GetServerInfo 获取服务器信息
func getServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		logrus.Errorf("获取服务器信息 InitCPU err:%v", err)
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		logrus.Errorf("获取服务器信息 InitRAM err:%v", err)
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		logrus.Errorf("获取服务器信息 InitDisk err:%v", err)
		return &s, err
	}
	return &s, nil
}

func (s *SystemApi) GetServerInfo(c *gin.Context) {
	if server, err := getServerInfo(); err != nil {
		logrus.Errorf("获取服务器信息 getServerInfo err:%v", err)
		global.Response(c, nil, err)
	} else {
		global.Response(c, gin.H{
			"server": server,
		}, nil)
	}
}
