package v1

import (
	"github.com/gin-gonic/gin"
	"tools/model/common/response"
	"tools/model/system/request"
	"tools/pkg/consts"
	"tools/service"
)

func ScanPort(ctx *gin.Context) {
	var target *request.PortScan
	if err := ctx.ShouldBind(&target); err != nil {
		response.FaithWithMessage(consts.PostBodyError, ctx)
		return
	}
	scanSrv := service.GetScanSrv()
	if resp, err := scanSrv.ScanPort(target); err != nil {
		response.FaithWithMessage(err.Error(), ctx)
		return
	} else {
		response.OkWithData(resp, ctx)
	}
}
