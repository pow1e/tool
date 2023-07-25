package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// RemoveSpace 去除字符串中的每个空格，获取到间隔空格的数据
func RemoveSpace(target string) []string {
	// 替换连续的多个空格为单个空格
	re := regexp.MustCompile(`\s+`)
	str := re.ReplaceAllString(target, " ")
	fmt.Println(str)

	// 使用空格分割字符串 获取元素
	data := strings.Split(str, " ")
	fmt.Println(data)
	return data
}
