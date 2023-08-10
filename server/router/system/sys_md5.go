package system

import (
	"github.com/gin-gonic/gin"
	v1 "tools/api/v1"
)

type Md5Router struct{}

func (m *Md5Router) InitMd5Router(router *gin.RouterGroup) {
	md5Router := router.Group("md5")
	{
		md5Router.POST("/encrypt", v1.EncryptMD5) // md5加密
		md5Router.POST("/decrypt", v1.DecryptMD5) // md5解密
	}
}
