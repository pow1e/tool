package response

import "tools/model/system"

type TaskListResp struct {
	Status   bool               `json:"status"`
	TaskList []*system.TaskList `json:"taskList"`
}
