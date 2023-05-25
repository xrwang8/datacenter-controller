package cmd

import (
	"datacenter-controller/mq"
	clientset "datacenter-controller/pkg/client/clientset/versioned"
	"flag"
	"path/filepath"

	"datacenter-controller/pkg/controller"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve",
	Long:  "",
	Run: func(c *cobra.Command, args []string) {
		klog.Info("start dc-controller")
		stopCh := make(chan struct{})
		mqClient, err := mq.NewMqClient("dc-controller")
		if err != nil {
			panic(err)
		}
		_, cfg, err := initClient()
		if err != nil {
			klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
		}
		dcClient, err := clientset.NewForConfig(cfg)
		if err != nil {
			klog.Fatalf("Error building dcClient clientset: %s", err.Error())
		}
		controller := controller.NewController(*dcClient, mqClient)
		err = controller.Subscription("volume-metrics")
		if err != nil {
			close(stopCh)
			klog.Errorf("%s topic not exist", "volume-metrics")
			return
		}

		<-stopCh
		klog.Errorf("shutdown")

	},
}

func initClient() (*kubernetes.Clientset, *rest.Config, error) {
	var err error
	var config *rest.Config
	// inCluster（Pod）、KubeConfig（kubectl）
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(可选) kubeconfig 文件的绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 文件的绝对路径")
	}
	flag.Parse()

	if config, err = rest.InClusterConfig(); err != nil {

		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}

	kubeclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, config, err
	}
	return kubeclient, config, nil
}
