package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	SUCCESS = 1
	ERROR   = 0
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func R(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(ctx *gin.Context) {
	R(SUCCESS, nil, "操作成功", ctx)
}

func OkWithMessage(message string, ctx *gin.Context) {
	R(SUCCESS, nil, message, ctx)
}

func OkWithData(data interface{}, ctx *gin.Context) {
	R(SUCCESS, data, "操作成功", ctx)
}

func OkWithDetailed(message string, data interface{}, ctx *gin.Context) {
	R(SUCCESS, data, message, ctx)
}

func Faith(ctx *gin.Context) {
	R(ERROR, nil, "操作失败", ctx)
}

func FaithWithMessage(message string, ctx *gin.Context) {
	R(ERROR, nil, message, ctx)
}

func FaithWithData(message string, data interface{}, ctx *gin.Context) {
	R(ERROR, data, message, ctx)
}
