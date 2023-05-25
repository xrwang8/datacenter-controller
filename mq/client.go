package mq

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// var MQClient *RocketmqService

func NewMqClient(group string) (*RocketmqService, error) {
	host := viper.GetString("MQ_SERVER")
	host = hostSchema(host)
	port := viper.GetString("MQ_PORT")

	endpoint := fmt.Sprintf("%s:%s", host, port)

	return NewRocketmqService([]string{endpoint}, group)

}

func hostSchema(host string) string {
	if !strings.HasPrefix(host, "http") {
		host = fmt.Sprintf("http://%s", host)
	}
	return host

}
