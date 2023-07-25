package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"tools/model/system"
	"tools/pkg/consts"
)

func UpdateSettings(target *system.Settings) error {
	dir, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprintf("获取当前根目录失败，err:%s", err))
	}
	// 将结构体反序列化
	data, err := json.Marshal(target)
	if err != nil {
		return errors.New(fmt.Sprintf("json序列化失败！err:%s", err))
	}
	err = ioutil.WriteFile(filepath.Join(dir, consts.SettingName), data, 0666)
	if err != nil {
		return errors.New(fmt.Sprintf("写入文件失败！err:%s", err))
	}
	return nil
}
