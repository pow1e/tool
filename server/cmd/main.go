package main

import (
	"go.uber.org/zap"
	"tools/core"
	"tools/global"
)

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @in                          header
// @BasePath                    /
func main() {
	global.Vp = core.InitViper() // 初始化Viper
	global.Log = core.InitZap()  // 初始化Zap
	zap.ReplaceGlobals(global.Log)
	core.InitRainbowTable()
	core.RunServer()
}
