package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DSN string
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
