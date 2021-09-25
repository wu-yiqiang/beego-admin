package utils

import (
	"fmt"
	"crypto/md5"
)

// MD5加密
func MD5(str string) string{
	md5str:=fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return md5str
}
