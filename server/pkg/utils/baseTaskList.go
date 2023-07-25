package utils

import (
	"strconv"
	"strings"
	"tools/global"
	"tools/model/system"
	"tools/model/system/response"
)

// BaseTaskList 传入终端命令运行结果 组装响应体 目前仅限于tasklist终端
func BaseTaskList(commandLine []byte) *response.TaskListResp {
	taskResp := new(response.TaskListResp)
	taskResp.Status = false
	lastEqIndex := strings.LastIndex(string(commandLine), "=")
	res := strings.TrimSpace(string(commandLine[lastEqIndex+1:]))
	data := RemoveSpace(res)

	global.Log.Info("运行命令行转换的字符是:" + res)

	imageName := data[0]
	pid := data[1]
	sessionName := data[2]
	session, _ := strconv.Atoi(data[3])
	memoryUsage := data[4] + data[5]
	taskList := system.NewTaskList(imageName, pid, sessionName, session, memoryUsage)

	taskResp.TaskList = append(taskResp.TaskList, taskList)
	taskResp.Status = true
	return taskResp
}
