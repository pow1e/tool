package admin

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

var (
	shell32       = syscall.NewLazyDLL("shell32.dll")
	shellExecuteW = shell32.NewProc("ShellExecuteW")
)

func Run(command string) ([]byte, error) {
	// 创建一个临时文件
	file, err := ioutil.TempFile("", "admin_*.bat")
	if err != nil {
		log.Fatalf("Couldn't create temporary file: %s", err.Error())
		return nil, err
	}

	// 获取文件的路径
	tempPath := file.Name()

	// 关闭临时文件
	err = file.Close()
	if err != nil {
		log.Fatalf("Couldn't close temporary file: %s", err.Error())
		return nil, err
	}

	// 将command写入文件中
	err = ioutil.WriteFile(tempPath, []byte(command), 0666)
	if err != nil {
		log.Fatalf("Couldn't write temporary file: %s", err.Error())
		return nil, err
	}

	// 将路径和其他参数转换为 Windows API 需要的类型
	lpOperation := syscall.StringToUTF16Ptr("runas")
	lpFile := syscall.StringToUTF16Ptr(tempPath)
	lpParameters := (*uint16)(unsafe.Pointer(nil))
	lpDirectory := (*uint16)(unsafe.Pointer(nil))
	nShowCmd := syscall.SW_HIDE

	// 调用 ShellExecuteW 函数以管理员身份运行脚本
	ret, _, err := shellExecuteW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(lpOperation)),
		uintptr(unsafe.Pointer(lpFile)),
		uintptr(unsafe.Pointer(lpParameters)),
		uintptr(unsafe.Pointer(lpDirectory)),
		uintptr(nShowCmd),
	)
	if ret <= 32 {
		return nil, err
	}

	// 执行命令并获取输出结果
	cmd := exec.Command("cmd", "/C", tempPath)

	output, _ := cmd.Output()

	// TODO 一定需要以管理员身份运行的是获取不到结果的
	//if err != nil {
	//	log.Fatalf("Couldn't run command: %s", err.Error())
	//	return nil, err
	//}

	// 将输出进行解码
	decoder := simplifiedchinese.GBK.NewDecoder()
	gbkOutPut, err := decoder.Bytes(output)
	if err != nil {
		log.Fatalf("Could not decode:%s\n", err.Error())
		return nil, err
	}

	// 删除临时文件
	err = os.Remove(tempPath)
	if err != nil {
		log.Printf("Couldn't delete temporary file: %s", err.Error())
		return nil, err
	}

	return gbkOutPut, nil
}
