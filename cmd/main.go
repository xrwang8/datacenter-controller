package main

import (
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"

	clientset "datacenter-controller/pkg/client/clientset/versioned"
	informers "datacenter-controller/pkg/client/informers/externalversions"
	conf "datacenter-controller/pkg/config"
)

var (
	onlyOneSignalHandler = make(chan struct{})
	shutdownSignals      = []os.Signal{os.Interrupt, syscall.SIGTERM}
)

// SetupSignalHandler 注册 SIGTERM 和 SIGINT 信号
// 返回一个 stop channel，该通道在捕获到第一个信号时被关闭
// 如果捕捉到第二个信号，程序将直接退出
func setupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler)

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)

	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // 第二个信号，直接退出
	}()

	return stop
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

func main() {
	flag.Parse()
	//设置一个信号处理，应用于优雅关闭
	stopCh := setupSignalHandler()

	_, cfg, err := initClient()
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	// 实例化一个 CronTab 的 ClientSet
	dcClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building crontab clientset: %s", err.Error())
	}

	dcInformerFactory := informers.NewSharedInformerFactory(dcClient, time.Second*30)

	// 启动 informer，开始List & Watch
	go dcInformerFactory.Start(stopCh)

}

func init() {
	cobra.OnInitialize(conf.InitConfig)
}
