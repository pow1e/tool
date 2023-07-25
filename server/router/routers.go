package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "tools/api/v1"
	"tools/middleware"
)

func Routers() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Cors()) 放行全部跨域请求
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	systemRouter := r.Group("api/v1")
	{
		// 查看mysql和redis是否启动
		check := systemRouter.Group("/check")
		{
			check.GET("", v1.CheckMysql)
			check.GET("/redis", v1.CheckRedis)
		}

		// 启动redis和mysql
		connect := systemRouter.Group("/connect")
		{
			connect.GET("/mysql", v1.MySqlConnect)
			connect.GET("/redis", v1.RedisConnect)
		}

		// md5接口
		md5 := systemRouter.Group("/md5")
		{
			// md5加密
			md5.POST("/encrypt", v1.EncryptMD5)
			// md5解密
			md5.POST("/decrypt", v1.DecryptMD5)
		}

		// 快捷指令
		shortcuts := systemRouter.Group("/shortcuts")
		shortcuts.Use(middleware.SettingExist())
		{
			shortcuts.POST("", v1.CreateShortcut)
			shortcuts.GET("", v1.GetShortcut)
			shortcuts.PUT("", v1.UpdateShortcut)
			shortcuts.GET("/run", v1.RunShortcut)
			shortcuts.DELETE("", v1.DeleteShortcut)
		}

		logs := systemRouter.Group("/log")
		{
			logs.DELETE("", v1.DeleteLogFile)
		}

	}
	return r
}
