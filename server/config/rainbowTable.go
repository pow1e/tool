package config

type RainbowRainbow struct {
	Status bool   `json:"status" mapstructure:"init-status" yaml:"init-status"`
	Path   string `json:"path" mapstructure:"rainbowTable-path" yaml:"rainbowTable-path"`
}
