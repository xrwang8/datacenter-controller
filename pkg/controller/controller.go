package controller

import (
	"context"
	"datacenter-controller/mq"
	clientset "datacenter-controller/pkg/client/clientset/versioned"
	"encoding/json"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
		for i := range msgs {
			klog.Infof("receive message:%v", string(msgs[i].Body))
			var volumeMetrics VolumeMetrics
			if err := json.Unmarshal(msgs[i].Body, &volumeMetrics); err != nil {
				klog.Errorf("unmarshal json to result error'%s'", err)
			}

			// 获取所有的自定义的cr信息
			dataCenter, err := c.clientset.DatacenterV1alpha1().DataCenters().Get(context.Background(), volumeMetrics.SubCenterId, v1.GetOptions{})
			if err != nil {
				klog.Errorf("Failed to get datacenter %v", err)
				return consumer.ConsumeSuccess, nil
			}
			memAllocatable := dataCenter.Spec.ResourceInfo.MemCapacity - volumeMetrics.MemAllocatable
			dataCenter.Spec.ResourceInfo.MemAllocatable = memAllocatable
			volumeCntAllocatable := dataCenter.Spec.ResourceInfo.VolumeCntCapacity - volumeMetrics.VolumeCntAllocatable
			dataCenter.Spec.ResourceInfo.VolumeCntAllocatable = volumeCntAllocatable
			dataCenter.Status.Idle.VolumeCntAllocatable = volumeCntAllocatable
			dataCenter.Status.Idle.MemAllocatable = memAllocatable
			c.clientset.DatacenterV1alpha1().DataCenters().Update(context.Background(), dataCenter, v1.UpdateOptions{})
			klog.Infof("update %v  datacenter metric %+v:", volumeMetrics.SubCenterId, dataCenter.Spec.ResourceInfo)
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
	VolumeCntAllocatable int    `json:"volumeCntAllocatable"`
	VolumeCntCapacity    int    `json:"volumeCntCapacity"`
}
