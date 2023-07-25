package core

import (
	"time"
	"tools/global"
	"tools/pkg/utils"
)

func InitRainbowTable() {
	if !global.Config.RainbowRainbow.Status {
		return
	}
	start := time.Now()
	err := utils.BuildRainbowTable()
	if err != nil {
		panic(err.Error())
		return
	}
	end := time.Now()
	elapsed := end.Sub(start) // 计算用时多久
	global.Log.Info("初始化彩虹表用时" + elapsed.String() + "秒")

	// 启动文件监视并接收错误信息
	errCh := make(chan error)
	go func() {
		errCh <- utils.WatchTextFile()
	}()
}
