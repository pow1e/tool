package service

import (
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strings"
	"tools/global"
	"tools/model/system"
	"tools/pkg/utils"
	"tools/pkg/utils/admin"
	"tools/pkg/utils/users"
)

func (s *ShortCutSrv) CreateShortCut(shortcut *system.Shortcut) error {
	// 校验参数是否合法
	err := shortcut.Validate(system.WithShorCutName(), system.WithShortCutDescription(),
		system.WithShortCutDescription(), system.WithShortCutType(), system.WithShortCutDirectives())
	if err != nil {
		return err
	}

	// 设置id
	shortcut.Id = uuid.New().String()

	// 赋值
	global.Settings.Shortcut = append(global.Settings.Shortcut, shortcut)

	return utils.UpdateSettings(global.Settings)
}

// GetShortCuts 获取快捷指令
func (s *ShortCutSrv) GetShortCuts(name string) (interface{}, error) {
	shortcuts := global.Settings.Shortcut
	if name == "" || len(name) == 0 {
		return shortcuts, nil
	}
	res := make([]*system.Shortcut, 0)
	for _, v := range shortcuts {
		if strings.Contains(v.Name, name) {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("暂无此快捷指令")
	}
	return res, nil
}

// UpdateShortcut 修改快捷指令信息
func (s *ShortCutSrv) UpdateShortcut(req *system.Shortcut) error {
	// 判断type是否合法
	if req.Type != "0" && req.Type != "1" {
		return errors.New("type应该为0或者1，请重试新输入！")
	}

	// 根据id查找是否存在需要求该的快捷指令
	for i, v := range global.Settings.Shortcut {
		if v.Id == req.Id {
			// 赋值
			global.Settings.Shortcut[i].Name = req.Name
			global.Settings.Shortcut[i].Description = req.Description
			global.Settings.Shortcut[i].Directives = req.Directives
			global.Settings.Shortcut[i].Type = req.Type
			err := utils.UpdateSettings(global.Settings)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("暂无此id，请重试")
}

// RunShortCut 运行终端命令
func (s *ShortCutSrv) RunShortCut(id string) (interface{}, error) {
	// 查询当前id是否存在
	var flag = false
	var target *system.Shortcut
	for _, v := range global.Settings.Shortcut {
		if v.Id == id {
			target = v
			flag = true
			break
		}
	}

	if !flag {
		return nil, errors.New("当前id不存在，请重试")
	}

	if len(target.Directives) == 0 {
		return nil, errors.New("指令不能为空！")
	}

	// 运行脚本
	run := make([]byte, 0)
	switch target.Type {
	case "0":
		res, err := users.Run(target.Directives)
		// 将运行的命令结果拼接到切片中
		if err != nil {
			return nil, err
		}
		run = append(run, res...)
		// TODO 美化返回
		//commandResp := utils.BaseTaskList(run)
		//return commandResp, nil
		return string(run), nil
	case "1":
		res, err := admin.Run(target.Directives)
		if err != nil {
			return nil, err
		}
		run = append(run, res...)
		// TODO 美化返回
		//commandResp := utils.BaseTaskList(run)
		//return commandResp, nil
		return string(run), nil
	default:
		return nil, errors.New("暂无此用户类型，请填写修改type为0或为1")
	}
}

// DeleteShortCut 根据id删除快捷方式
func (s *ShortCutSrv) DeleteShortCut(ids []string) error {
	var flag = false
	for i := len(global.Settings.Shortcut) - 1; i >= 0; i-- {
		elem := global.Settings.Shortcut[i].Id
		for _, deleteId := range ids {
			if elem == deleteId {
				// 执行删除功能
				global.Settings.Shortcut = append(global.Settings.Shortcut[:i], global.Settings.Shortcut[i+1:]...)
				flag = true
				break // 找到并删除元素后，退出内部循环。
			}
		}
	}

	if !flag {
		return errors.New("暂无此id，请重试！")
	}
	// 更新文件操作
	err := utils.UpdateSettings(global.Settings)
	if err != nil {
		global.Log.Error("更新文件失败", zap.Error(err))
		return errors.New("更新setting.json文件出错，请重试！")
	}

	return nil
}
