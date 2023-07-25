package users

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"tools/global"
)

// Run 传入执行命令 返回cmd运行结果
func Run(command string) ([]byte, error) {
	// 创建临时文件
	uuidStr := uuid.NewString()
	file, err := ioutil.TempFile("", uuidStr+"*.bat")
	if err != nil {
		global.Log.Error("创建临时文件失败", zap.Error(err))
		return nil, err
	}

	// 获取文件路径
	tmpFilePath := file.Name()

	// 关闭临时文件 关闭临时文件只是释放了文件的占用，以便其他进程或操作可以对该文件访问
	err = file.Close()
	if err != nil {
		global.Log.Error("文件打开失败", zap.Error(err))
		return nil, err
	}

	// 将操作写进文件中
	err = ioutil.WriteFile(tmpFilePath, []byte(command), 0664)
	if err != nil {
		global.Log.Error("读取文件失败", zap.Error(err))
		return nil, err
	}

	// 运行临时文件
	cmd := exec.Command("cmd", "/C", tmpFilePath)

	output, err := cmd.Output()
	if err != nil {
		global.Log.Error("命令输出错误", zap.Error(err))
		return nil, err
	}

	// 将输出进行解码
	decoder := simplifiedchinese.GBK.NewDecoder()
	gbkOutPut, err := decoder.Bytes(output)
	if err != nil {
		global.Log.Error("Could not decode", zap.Error(err))
		return nil, err
	}

	// 删除临时文件
	err = os.Remove(tmpFilePath)
	if err != nil {
		log.Printf("删除临时文件出错：%v", err)
		return nil, err
	}

	return gbkOutPut, nil
}
