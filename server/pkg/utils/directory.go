package utils

import (
	"errors"
	"os"
)

// PathExists 判断当前文件夹是否存在
func PathExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	// 不存在错误
	if err == nil {
		if fileInfo.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
