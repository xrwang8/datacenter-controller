package cmd

import (
	"datacenter-controller/mq"
	"datacenter-controller/pkg/config"

	"github.com/spf13/cobra"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var ServeCmd = &cobra.Command{
	Use:   "dc-controller",
	Short: "dc-controller",
	Long:  "",
	Run: func(c *cobra.Command, args []string) {
		klog.Info("start dc-controller")
		stopCh := signals.SetupSignalHandler()

		conf := config.NewConfig()
		klog.Infof("config:%v", *conf)
		mqClient, err := mq.NewMqClient("dc-controller")
		if err != nil {
			panic(err)
		}

	},
}
