package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"tools/config"
	"tools/model/system"
)

var (
	Vp           *viper.Viper
	Config       config.Config
	Log          *zap.Logger
	Settings     *system.Settings
	RainBowTable *map[string]string
)
