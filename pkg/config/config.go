package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Variables *viper.Viper
}

func NewConfig() *Config {
	viper.AutomaticEnv()
	return &Config{
		Variables: viper.GetViper(),
	}
}
