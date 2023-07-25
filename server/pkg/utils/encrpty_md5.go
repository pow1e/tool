package utils

import (
	"crypto/md5"
	"encoding/base32"
	"fmt"
	"strings"
)

// GetMD5 加密md5 传入一个明文以及进制数
func GetMD5(plaintext string) []string {
	encrypts := make([]string, 0)
	bytes := []byte(plaintext)
	hash := md5.Sum(bytes)
	base16Lowercase := fmt.Sprintf("%x", hash)          // 16进制小写
	base16Uppercase := strings.ToUpper(base16Lowercase) // 16进制大写
	base32Lowercase := convertToBase32(hash)            // 32进制小写
	base32Uppercase := strings.ToUpper(base32Lowercase) // 32进制大写
	encrypts = append(encrypts, base16Lowercase, base16Uppercase, base32Lowercase, base32Uppercase)
	return encrypts
}

// 将MD5哈希转换为32进制
func convertToBase32(hash [16]byte) string {
	// 指定32进制字符集，设置无填充模式
	encoding := base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").WithPadding(base32.NoPadding)
	result := make([]byte, encoding.EncodedLen(len(hash)))
	encoding.Encode(result, hash[:])
	return string(result)
}
