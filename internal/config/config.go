package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiKey string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	AppConfig = &Config{
		ApiKey: viper.GetString("KUBARCLOUD_APIKEY"),
	}
}
