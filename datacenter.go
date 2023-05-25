package main

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"

	"datacenter-controller/cmd"
	conf "datacenter-controller/pkg/config"
)

func init() {
	cobra.OnInitialize(conf.InitConfig)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
