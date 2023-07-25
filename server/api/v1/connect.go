package v1

import (
	"github.com/gin-gonic/gin"
	"tools/global"
	"tools/model/common/response"
	"tools/service"
)

// MySqlConnect 启动mysql服务
func MySqlConnect(ctx *gin.Context) {
	srv := service.GetSysSrv()
	if err := srv.StartMysqlService(); err != nil {
		global.Log.Error(err.Error())
		response.FaithWithMessage("启动失败", ctx)
		return
	}
	response.Ok(ctx)
}

// RedisConnect 启动redis服务
func RedisConnect(ctx *gin.Context) {
	srv := service.GetSysSrv()
	err := srv.StartRedisService()
	if err != nil {
		global.Log.Error(err.Error())
		response.FaithWithMessage("启动redis失败，请重试", ctx)
		return
	}
	response.Ok(ctx)
}
