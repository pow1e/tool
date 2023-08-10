package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// GetMD5 加密md5 传入一个明文以及进制数
func GetMD5(plaintext string) []string {
	encrypts := make([]string, 0)
	bytes := []byte(plaintext)
	hash := md5.Sum(bytes)

	md5Hash32 := hex.EncodeToString(hash[:])

	md5Hash32Upper := strings.ToUpper(md5Hash32)

	md5Hash16 := md5Hash32[8:24]

	md5Hash16Upper := strings.ToUpper(md5Hash16)

	encrypts = append(encrypts, md5Hash16, md5Hash16Upper, md5Hash32, md5Hash32Upper)
	return encrypts
}
