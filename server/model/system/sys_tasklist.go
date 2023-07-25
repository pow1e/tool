package system

type TaskList struct {
	ImageName   string `json:"image"`
	PID         string `json:"pid"`
	SessionName string `json:"sessionName"`
	Session     int    `json:"session"`
	MemoryUsage string `json:"memoryUsage"`
}

func NewTaskList(imageName string, pid string, sessionName string, session int, memoryUsage string) *TaskList {
	return &TaskList{
		ImageName:   imageName,
		PID:         pid,
		SessionName: sessionName,
		Session:     session,
		MemoryUsage: memoryUsage,
	}
}
