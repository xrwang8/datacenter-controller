package main

import (
	"datacenter-controller/mq"
	informers "datacenter-controller/pkg/client/informers/externalversions/datacenter.pcl.ac.cn/v1alpha1"
)

type Controller struct {
	informer informers.DataCenterInformer
	mqClient *mq.RocketmqService
}

func NewController(informer informers.DataCenterInformer, mqClient *mq.RocketmqService) *Controller {
	//使用client 和前面创建的 Informer，初始化了自定义控制器
	controller := &Controller{
		informer: informer,
		mqClient: mqClient,
	}

	return controller
}
