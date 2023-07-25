package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"tools/global"
	"tools/pkg/consts"
	"tools/pkg/utils"
	"tools/pkg/window"
)

func InitViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取文件失败！: %s \n", err))
	}
	// 监视当前文件是否修改
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			err = window.ErrorWindow(window.ErrorCommon())
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(err.Error())
			panic(err.Error())
		}
	})

	// 赋值
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	// 读取用户信息文件
	ReadSettings(v)
	return v
}

func ReadSettings(v *viper.Viper) {

	v.SetConfigFile(consts.SettingName)

	// 启动监听
	v.WatchConfig()

	// 读取文件
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	// 监听配置文件的变化
	v.OnConfigChange(func(e fsnotify.Event) {
		//fmt.Printf("%s文件发生改变\n", e.Name)

		// 清空数据
		global.Settings = nil

		// 文件发生变化时重新读取配置
		if err = v.ReadInConfig(); err != nil {
			fmt.Println("Failed to read config file:", err)
			panic(err.Error())
		}

		// 重新反序列化
		if err = v.Unmarshal(&global.Settings); err != nil {
			fmt.Println("Failed to re-serialize:", err)
			panic(err.Error())
		}

		// 判断当前id是否存在为空的
		flag, err := utils.AutoFillId(global.Settings)
		if err != nil {
			panic(err.Error())
		}
		// 如果保存后仍然存在id为空情况下，则重新写入文件
		if flag {
			err := utils.UpdateSettings(global.Settings)
			if err != nil {
				panic(err.Error())
			}
			return
		}
	})

	// 序列化
	if err = v.Unmarshal(&global.Settings); err != nil {
		fmt.Println("Failed to re-serialize:", err)
		panic(err.Error())
	}

	// 查看当前id是否为空 如果为空则自动赋值
	flag, err := utils.AutoFillId(global.Settings)
	if err != nil {
		panic(err.Error())
	}

	// id 为空情况下
	if flag {
		// 重新写入文件
		err = utils.UpdateSettings(global.Settings)
		if err != nil {
			panic(err.Error())
		}
	}

}
