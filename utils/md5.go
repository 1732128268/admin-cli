package utils

import (
	"admin-cli/config"
	"crypto/md5"
	"encoding/hex"
)

func GenMd5(src string) string {
	conf := config.GetConfig()
	srcBytes := []byte(src + conf.Salt)

	md5New := md5.New()

	md5Bytes := md5New.Sum(srcBytes)

	md5String := hex.EncodeToString(md5Bytes)
	return md5String

}
