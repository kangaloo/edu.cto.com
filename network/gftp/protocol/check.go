package protocol

import (
	"crypto/md5"
	"encoding/hex"
)

// 利用闭包函数增量的检查文件的 md5 值
func CheckSum() func([]byte) string {

	checker := md5.New()

	Checker := func(b []byte) string {
		checker.Write(b)
		sum := checker.Sum(nil)
		return hex.EncodeToString(sum)
	}
	return Checker

}
