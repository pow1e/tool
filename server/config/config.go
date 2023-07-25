package config

type Config struct {
	Zap            *Zap            `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis          *Redis          `mapstructure:"redis" json:"redis" yaml:"redis"`
	System         *System         `mapstructure:"system" json:"system" yaml:"system"`
	Mysql          *Mysql          `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	RainbowRainbow *RainbowRainbow `mapstructure:"rainbow-table" json:"rainbow-table" yaml:"rainbow-table"`
}
