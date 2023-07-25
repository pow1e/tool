package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"tools/core/internal"
	"tools/global"
	"tools/pkg/utils"
)

func InitZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	// 检查日志并且删除
	//checkAndDeleteOldLogs(global.Config.Zap.Director, time.Hour*24)
	return logger
}

//func checkAndDeleteOldLogs(logDirectory string, retentionPeriod time.Duration) {
//	keepTime := time.Now().Add(-retentionPeriod)
//	files, err := os.ReadDir(logDirectory)
//	if err != nil {
//		log.Printf("failed to read log directory: %v", err)
//		return
//	}
//
//	for _, file := range files {
//		if !file.IsDir() {
//			filePath := path.Join(logDirectory, file.Name())
//			fileInfo, err := os.Stat(filePath)
//			if err != nil {
//				log.Printf("failed to get file info: %s, %v", file.Name(), err)
//				continue
//			}
//
//			if fileInfo.ModTime().Before(keepTime) {
//				if err := os.Remove(filePath); err != nil {
//					log.Printf("failed to delete log file: %s, %v", file.Name(), err)
//				} else {
//					log.Printf("deleted log file: %s", file.Name())
//				}
//			}
//		}
//	}
//}
