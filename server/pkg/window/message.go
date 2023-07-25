package window

import (
	"fmt"
	"syscall"
	"tools/pkg/consts"
	"unsafe"
)

func MessageWindow(title, content string) error {
	var flags uint = consts.ButtonOkcancel | consts.ButtonIconwarning
	ret, _, err := messageBox.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(content))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags),
	)
	if err != nil {
		return err
	}
	switch int32(ret) {
	case 1:
		fmt.Println("用户点击了确认")
	case 2:
		fmt.Println("用户点击了取消")
	default:
		fmt.Println("用户关闭了弹窗")
	}
	return nil
}
