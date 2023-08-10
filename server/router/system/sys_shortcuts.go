package system

import (
	"github.com/gin-gonic/gin"
	v1 "tools/api/v1"
	"tools/middleware"
)

type ShortCutRouter struct{}

func (s *ShortCutRouter) InitShorCutRouter(router *gin.RouterGroup) {
	shortcutsRouter := router.Group("shortcuts").Use(middleware.SettingExist())
	{
		shortcutsRouter.POST("", v1.CreateShortcut)
		shortcutsRouter.GET("", v1.GetShortcut)
		shortcutsRouter.PUT("", v1.UpdateShortcut)
		shortcutsRouter.GET("/run", v1.RunShortcut)
		shortcutsRouter.DELETE("", v1.DeleteShortcut)
	}
}
