package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tools/global"
	"tools/model/common/request"
	"tools/model/common/response"
	"tools/pkg/consts"
)

func DeleteLogFile(ctx *gin.Context) {
	var req *request.IdsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		global.Log.Error("绑定参数失败", zap.Error(err))
		response.FaithWithMessage(consts.PostBodyError, ctx)
		return
	}
	if len(req.Ids) == 0 {
		response.FaithWithMessage("请正确填写所需要删除的日志格式", ctx)
		return
	}

}
