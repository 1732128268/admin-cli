package captcha

import (
	"admin-cli/global"
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
	"time"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 60 * 5,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := global.Redis.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		logrus.Errorf("set redis error: %v", err)
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.Redis.Get(rs.Context, key).Result()
	if err != nil {
		logrus.Errorf("get redis error: %v", err)
		return ""
	}
	if clear {
		err := global.Redis.Del(rs.Context, key).Err()
		if err != nil {
			logrus.Errorf("del redis error: %v", err)
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
