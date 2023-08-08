package system

import (
	"github.com/gin-gonic/gin"
	v1 "tools/api/v1"
)

type ScanRouter struct {
}

func (s ScanRouter) InitScan(router *gin.RouterGroup) {
	scanGroup := router.Group("scan")
	{
		scanGroup.POST("/port", v1.ScanPort)
	}
}
