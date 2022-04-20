package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMd5(src string) string {
	srcBytes := []byte(src)

	md5New := md5.New()

	md5New.Write(srcBytes)

	md5String := hex.EncodeToString(md5New.Sum(nil))
	return md5String
}
