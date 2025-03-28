package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5 定一个函数传入字符串，返回md5加密后的字符串
func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
