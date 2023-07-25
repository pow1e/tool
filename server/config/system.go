package config

type System struct {
	Env          string `mapstructure:"env" yaml:"env" json:"env"`
	Addr         string `mapstructure:"addr" yaml:"addr" json:"addr"`
	RouterPrefix string `mapstructure:"router-prefix" yaml:"router-prefix" json:"router-prefix"`
}
