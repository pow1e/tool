package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"tools/global"
	"tools/model/common/response"
	"tools/pkg/consts"
)

// SettingExist 判断setting.json文件是否存在 不存在则会创建一个新文件，存在则放行
func SettingExist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := os.Stat(consts.SettingName)
		if err != nil {
			if os.IsNotExist(err) {
				// 当前文件不存在
				global.Log.Info("当前文件不存在，在创建新文件")

				dir, err := os.Getwd()
				if err != nil {
					response.FaithWithMessage("获取当前文件目录失败！", ctx)
					global.Log.Error(err.Error())
					ctx.Abort()
					return
				}

				filePath := filepath.Join(dir, consts.SettingName)
				file, err := os.Create(filePath)
				if err != nil {
					response.FaithWithMessage("用户文件不存在，创建文件失败！", ctx)
					global.Log.Error(err.Error())
					ctx.Abort()
					return
				}
				file.Close()
				return
			} else {
				response.FaithWithMessage("无法访问用户文件信息！", ctx)
				global.Log.Error(err.Error())
				ctx.Abort()
				return
			}
		}
	}
}
