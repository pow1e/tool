package initialize

import (
	"github.com/gin-gonic/gin"
	"tools/global"
	"tools/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Cors()) 放行全部跨域请求

	ApiGroupEnter := router.RouterGroupEnter
	DefaultPrefix := r.Group(global.Config.System.RouterPrefix)
	{
		ApiGroupEnter.InitShorCutRouter(DefaultPrefix)
		ApiGroupEnter.InitMd5Router(DefaultPrefix)
	}
	return r
}
