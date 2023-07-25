package service

import (
	"strconv"
	"strings"
	"tools/global"
	"tools/model/system"
	"tools/model/system/response"
	"tools/pkg/utils"
	"tools/pkg/utils/users"
)

func (s *SysSrv) StartMysqlService() error {
	return nil
}

func (s *SysSrv) StartRedisService() error {
	return nil
}

func (s *SysSrv) CheckMysql() (interface{}, error) {
	taskResp := new(response.TaskListResp)
	taskResp.Status = false
	byteRes, err := users.Run("tasklist /FI \"IMAGENAME eq mysqld.exe\"\n")
	if err != nil {
		return taskResp, err
	}

	lastEqIndex := strings.LastIndex(string(byteRes), "=")
	res := strings.TrimSpace(string(byteRes[lastEqIndex+1:]))
	data := utils.RemoveSpace(res)

	global.Log.Info("运行命令行转换的字符是:" + res)

	imageName := data[0]
	pid := data[1]
	sessionName := data[2]
	session, _ := strconv.Atoi(data[3])
	memoryUsage := data[4] + data[5]
	taskList := system.NewTaskList(imageName, pid, sessionName, session, memoryUsage)

	taskResp.TaskList = append(taskResp.TaskList, taskList)
	taskResp.Status = true

	return taskResp, nil
}

func (s *SysSrv) CheckRedis() (interface{}, error) {
	taskResp := new(response.TaskListResp)
	taskResp.Status = false
	byteRes, err := users.Run("tasklist /FI \"IMAGENAME eq redis-server.exe\"\n")
	if err != nil {
		return taskResp, err
	}

	lastEqIndex := strings.LastIndex(string(byteRes), "=")
	res := strings.TrimSpace(string(byteRes[lastEqIndex+1:]))
	data := utils.RemoveSpace(res)

	global.Log.Info("运行命令行转换的字符是:" + res)

	imageName := data[0]
	pid := data[1]
	sessionName := data[2]
	session, _ := strconv.Atoi(data[3])
	memoryUsage := data[4] + data[5]
	taskList := system.NewTaskList(imageName, pid, sessionName, session, memoryUsage)

	taskResp.TaskList = append(taskResp.TaskList, taskList)
	taskResp.Status = true

	return taskResp, nil
}
