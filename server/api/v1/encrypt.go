package v1

import (
	"github.com/gin-gonic/gin"
	"tools/model/common/response"
	"tools/model/system"
	"tools/pkg/consts"
	"tools/service"
)

func Encrypt(ctx *gin.Context) {
	var req *system.AlgorithmReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FaithWithMessage(consts.PostBodyError, ctx)
		return
	}
	encrypts, err := service.GetAlgorithmSrv().Encrypt(req)
	if err != nil {
		response.FaithWithMessage(err.Error(), ctx)
	}
	response.OkWithDetailed("加密成功，注意aes加密的key长度是16位，不够默认以0填充", encrypts, ctx)
}
