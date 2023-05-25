package config

import (
	"sync"

	"github.com/spf13/viper"
)

func InitConfig() {

	viper.SetDefault("MQ_SERVER", "10.102.10.1")
	viper.SetDefault("MQ_PORT", "9876")

	viper.AutomaticEnv()
}

type MqConfig struct {
	Server string `json:"server,omitempty" yaml:"server"`
	Port   string `json:"port,omitempty" yaml:"port"`
	Key    string `json:"key,omitempty" yaml:"key"`
	Secret string `json:"passwd,omitempty" yaml:"passwd"`
}

type Config struct {
	Mq *MqConfig `json:"mq,omitempty"`
}

var conf *Config

func NewConfig() *Config {

	var once sync.Once
	once.Do(func() {
		conf = &Config{
			Mq: &MqConfig{
				Server: viper.GetString("MQ_SERVER"),
				Port:   viper.GetString("MQ_PORT"),
				Key:    viper.GetString("MQ_KEY"),
				Secret: viper.GetString("MQ_SECRET"),
			},
		}
	})
	return conf
}
