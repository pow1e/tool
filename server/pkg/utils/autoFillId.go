package utils

import (
	"github.com/google/uuid"
	"tools/model/system"
)

// AutoFillId 保存时候自动修改id为空的字段 存在id为空返回true 否则返回false
func AutoFillId(settings *system.Settings) (bool, error) {
	// 快捷指令id自动填充
	var flag = false
	shortcuts := settings.Shortcut
	for i, v := range shortcuts {
		if len(v.Id) == 0 || v.Id == "" {
			shortcuts[i].Id = uuid.New().String()
			flag = true
		}
	}

	// TODO 自动填充id其他设置

	return flag, nil
}
