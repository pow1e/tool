package main

import (
	"go.uber.org/zap"
	"tools/core"
	"tools/global"
	"tools/model/system"
)

func main() {
	global.Vp = core.InitViper() // 初始化Viper
	global.Log = core.InitZap()  // 初始化Zap
	zap.ReplaceGlobals(global.Log)
	core.InitRainbowTable()
	system.InitAlgorithmMap()
	core.RunServer()
}
