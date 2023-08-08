package router

import "tools/router/system"

type RouterGroup struct {
	system.ShortCutRouter
	system.Md5Router
	system.ScanRouter
	system.AlgorithmRouter
}

var RouterGroupEnter = new(RouterGroup)
