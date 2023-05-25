package controller

import (
	"context"
	"datacenter-controller/mq"
	clientset "datacenter-controller/pkg/client/clientset/versioned"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"k8s.io/klog"
)

type Controller struct {
	clientset clientset.Clientset
	mqClient  *mq.RocketmqService
}

func NewController(clientset clientset.Clientset, mqClient *mq.RocketmqService) *Controller {
	controller := &Controller{
		clientset: clientset,
		mqClient:  mqClient,
	}
	return controller
}

func (c *Controller) Subscription(topic string) error {
	err := c.mqClient.Consumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			klog.Infof("receive message:%v", string(msg.Body))
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		klog.V(4).Infof("from sub cluster subscribe info error'%s'", err)
		return err
	}
	err = c.mqClient.Consumer.Start()
	if err != nil {
		klog.V(4).Infof(" subscribe info error'%s'", err)
		return err
	}
	return nil
}

type VolumeMetrics struct {
	SubCenterId          string `json:"subCenterId"`
	MemAllocatable       int64  `json:"memAllocatable"`
	MemCapacity          int64  `json:"memCapacity"`
	VolumeCntAllocatable int64  `json:"volumeCntAllocatable"`
	VolumeCntCapacity    int64  `json:"volumeCntCapacity"`
}
