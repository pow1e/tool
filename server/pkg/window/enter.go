package window

import "syscall"

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	messageBox          = user32.NewProc("MessageBoxW")
	getActiveWindow     = user32.NewProc("GetActiveWindow")
	setForegroundWindow = user32.NewProc("SetForegroundWindow")
)

func ErrorCommon() (string, string) {
	return "错误", "内部错误，请重试！"
}
