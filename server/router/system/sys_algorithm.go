package system

import (
	"github.com/gin-gonic/gin"
	v1 "tools/api/v1"
)

type AlgorithmRouter struct{}

func (a *AlgorithmRouter) InitAlgorithmRouter(router *gin.RouterGroup) {
	router.POST("/encrypt", v1.Encrypt)
	router.POST("/decrypt", func(ctx *gin.Context) {
	})
}
