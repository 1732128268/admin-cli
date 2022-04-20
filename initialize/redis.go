package initialize

import (
	"admin-cli/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// InitRedis 初始化redis
func InitRedis() error {
	redisConfig := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
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
