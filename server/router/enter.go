package router

import "tools/router/system"

type RouterGroup struct {
	system.ShortCutRouter
	system.Md5Router
}

var RouterGroupEnter = new(RouterGroup)
