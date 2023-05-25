package config

import (
	"github.com/spf13/viper"
)

func InitConfig() {

	viper.SetDefault("MQ_SERVER", "10.0.102.10")
	viper.SetDefault("MQ_PORT", "9876")

	viper.AutomaticEnv()
}
