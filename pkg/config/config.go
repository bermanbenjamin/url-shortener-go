package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Variables *viper.Viper
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	uri := viper.GetString("database.uri")

	log.Println("Connect to uri:" + uri)

	return &Config{
		Variables: viper.GetViper(),
	}
}
