package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tools/global"
	"tools/model/common/request"
	"tools/model/common/response"
	"tools/model/system"
	"tools/pkg/consts"
	"tools/service"
)

// CreateShortcut 创建快捷指令
func CreateShortcut(ctx *gin.Context) {
	var req *system.Shortcut
	if err := ctx.ShouldBind(&req); err != nil {
		global.Log.Error("绑定参数失败", zap.Error(err))
		response.FaithWithMessage("绑定参数失败，请重试", ctx)
		return
	}

	srv := service.GetShortCutSrv()
	err := srv.CreateShortCut(req)
	if err != nil {
		global.Log.Error("创建快捷指令失败", zap.Error(err))
		response.FaithWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

// GetShortcut 获取快捷指令
func GetShortcut(ctx *gin.Context) {
	name := ctx.Query("name")
	global.Log.Info("获取快捷指令:" + name)
	srv := service.GetShortCutSrv()
	cuts, err := srv.GetShortCuts(name)
	if err != nil {
		global.Log.Error("获取快捷直径失败!", zap.Error(err))
		response.FaithWithMessage("暂时无此快捷指令！", ctx)
		return
	}
	response.OkWithData(cuts, ctx)
}

// UpdateShortcut 更新快捷指令
func UpdateShortcut(ctx *gin.Context) {
	var req *system.Shortcut
	if err := ctx.ShouldBind(&req); err != nil {
		global.Log.Error("获取参数失败", zap.Error(err))
		response.FaithWithMessage("获取参数失败！", ctx)
		return
	}
	global.Log.Info(fmt.Sprintf("修改快捷指令：%v\n", *req))
	srv := service.GetShortCutSrv()
	err := srv.UpdateShortcut(req)
	if err != nil {
		global.Log.Error("更新失败！", zap.Error(err))
		response.FaithWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("更新成功！", ctx)
}

// RunShortcut 运行快捷指令
func RunShortcut(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		response.FaithWithMessage(consts.GetQueryError, ctx)
		return
	}
	srv := service.GetShortCutSrv()
	commandResp, err := srv.RunShortCut(id)
	if err != nil {
		global.Log.Error("运行失败！", zap.Error(err))
		response.FaithWithMessage(err.Error(), ctx)
		return
	}
	// TODO 美化返回值 不需要断言
	if data, ok := commandResp.(string); !ok {
		response.FaithWithMessage("数据获取失败，请重试", ctx)
	} else {
		response.OkWithData(data, ctx)
	}

}

// DeleteShortcut 根据id删除快捷方式
func DeleteShortcut(ctx *gin.Context) {
	var req *request.IdsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		global.Log.Error("绑定参数失败", zap.Error(err))
		response.FaithWithMessage(consts.PostBodyError, ctx)
		return
	}
	if len(req.Ids) == 0 {
		response.FaithWithMessage("请正确填写所需要删除的id", ctx)
		return
	}

	srv := service.GetShortCutSrv()
	if err := srv.DeleteShortCut(req.Ids); err != nil {
		response.FaithWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}
