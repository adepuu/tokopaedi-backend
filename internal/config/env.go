package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(".env")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return config
}
