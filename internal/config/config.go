package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() *viper.Viper {
	viper.SetConfigName("env")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.WatchConfig()

	return viper.GetViper()
}
