package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"tools/global"
	"tools/model/common/response"
	"tools/model/system/request"
	resp "tools/model/system/response"
	"tools/pkg/consts"
	"tools/pkg/utils"
)

// EncryptMD5 md5加密
func EncryptMD5(ctx *gin.Context) {
	var req *request.Md5Encrypt
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FaithWithMessage("绑定参数失败", ctx)
		return
	}
	if req.Plaintext == "" {
		response.FaithWithMessage(consts.PostBodyError, ctx)
		return
	}
	data := utils.GetMD5(req.Plaintext)
	md5Resp := resp.MD5Resp{
		Base16Lowercase: data[0],
		Base16Uppercase: data[1],
		Base32Lowercase: data[2],
		Base32Uppercase: data[3],
	}
	response.OkWithDetailed("加密成功", md5Resp, ctx)
}

// DecryptMD5 md5解密
func DecryptMD5(ctx *gin.Context) {
	var req *request.Md5Decrypt
	if err := ctx.ShouldBind(&req); err != nil {
		response.FaithWithMessage("绑定参数失败", ctx)
		return
	}
	if len(req.Ciphertext) != 16 && len(req.Ciphertext) != 32 {
		response.FaithWithMessage("请正确填写md5密文(16位或32位)！", ctx)
		return
	}

	if global.RainBowTable == nil {
		var err error
		err = utils.BuildRainbowTable()
		if err != nil {
			global.Log.Error(fmt.Sprintf("Error buildingRainbowTable: %v", err))
			response.FaithWithMessage("构建彩虹表失败！", ctx)
			return
		}
	}

	// 转换成小写
	req.Ciphertext = strings.ToLower(req.Ciphertext)

	var plaintext string

	switch len(req.Ciphertext) {
	case 16:
		for key, value := range *global.RainBowTable {
			if key[8:24] == req.Ciphertext {
				plaintext = value
				break
			}
		}
	case 32:
		// 默认是小数32位置
		plaintext = utils.LookupRainbowTable(strings.ToLower(req.Ciphertext), *global.RainBowTable)
	}

	if len(plaintext) == 0 || plaintext == "" {
		response.FaithWithMessage("解密失败", ctx)
		return
	}

	response.OkWithDetailed("解密成功", plaintext, ctx)
}
