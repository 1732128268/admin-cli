package initialize

import (
	"admin-cli/config"
	"admin-cli/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// InitRedis 初始化redis
func InitRedis() error {
	configData := config.GetConfig()
	redisConfig := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configData.Redis.Host, configData.Redis.Port),
		Password: configData.Redis.Password,
		DB:       configData.Redis.DB,
	}
	client := redis.NewClient(redisConfig)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("redis connect error %v", err)
		return err
	}
	global.Redis = client
	return nil
}
