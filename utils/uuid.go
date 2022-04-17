package utils

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

func GetUid() string {
	uid, _ := uuid.NewV4()
	return uuid.Must(uid, nil).String()
}

// GetRandomNum 获取随机8位数字
func GetRandomNum() string {
	return GetUid()[:8]
}

// GetOrderId 获取订单Id
func GetOrderId() string {
	return time.Now().Format("20060102150405") + GetRandomNum()
}
