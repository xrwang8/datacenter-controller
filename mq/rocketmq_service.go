package mq

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
	"os"
	"strings"
	"sync"
	"time"
)

var doOnce sync.Once

type RocketmqService struct {
	Producer rocketmq.Producer
	Consumer rocketmq.PushConsumer
}

type IMqService interface {
	ProduceMsg(topic string, data interface{}) error
	ProduceMsgWithTag(topic, tag string, data interface{}) error
	ConsumeMsg(topic string)
	ConsumeMsgWithTag(topic, tag string)
	Shutdown()
}

func NewRocketmqService(endpoints []string, group string) (*RocketmqService, error) {
	rlog.SetLogLevel("warn")

	p, err := NewProducer(endpoints, group)
	if err != nil {
		return nil, err
	}

	c, err := NewPushConsumer(endpoints, group)
	if err != nil {
		return nil, err
	}

	var rockermqService *RocketmqService
	doOnce.Do(func() {
		rockermqService = &RocketmqService{
			Producer: p,
			Consumer: c,
		}
	})
	return rockermqService, nil
}

func NewProducer(endpoints []string, group string) (rocketmq.Producer, error) {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(endpoints),
		producer.WithGroupName(group),
		producer.WithRetry(5),
		producer.WithSendMsgTimeout(3*time.Second),
		producer.WithQueueSelector(producer.NewHashQueueSelector()),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: viper.GetString("MQ_KEY"),
			SecretKey: viper.GetString("MQ_SECRET"),
			// SecurityToken: "",
		}),
	)

	if err != nil {
		klog.Errorf("new producer error: %s", err.Error())
		return nil, err
	}

	err = p.Start()
	if err != nil {
		klog.Errorf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	return p, nil
}
func NewPushConsumer(endpoints []string, group string) (rocketmq.PushConsumer, error) {

	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNameServer(endpoints),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: viper.GetString("MQ_KEY"),
			SecretKey: viper.GetString("MQ_SECRET"),
			// SecurityToken: "",
		}),
	)
	if err != nil {
		klog.Errorf("new consumer errror:%s", err.Error())
		return nil, err
	}
	return c, nil
}

func (c *RocketmqService) Shutdown() {
	if c.Producer != nil {
		err := c.Producer.Shutdown()
		if err != nil {
			klog.Errorf("shutdown producer error: %s", err.Error())
		}
	}
}

func (c *RocketmqService) ProduceMsg(topic string, data interface{}) error {
	msg := primitive.NewMessage(topic, Marshal(data))
	res, err := c.Producer.SendSync(context.Background(), msg)
	if err != nil {
		klog.Errorf("Send message error: %s\n", err)
		return err
	}

	klog.V(5).Infof("Send message success: result=%s\n", res.String())
	return nil
}

func (c *RocketmqService) ProduceMsgWithShardingKey(topic string, data interface{}, shardingKey string) error {
	msg := primitive.NewMessage(topic, Marshal(data))
	msg.WithShardingKey(shardingKey)
	res, err := c.Producer.SendSync(context.Background(), msg)
	if err != nil {
		klog.Errorf("Send message error: %s\n", err)
		return err
	}

	klog.V(5).Infof("Send message success: result=%s\n", res.String())
	return nil
}
func (c *RocketmqService) ProduceMsgAndKeyWithShardingKey(topic, tag string, data interface{}, shardingKey string) error {
	msg := primitive.NewMessage(topic, Marshal(data))
	msg.WithShardingKey(shardingKey)
	msg.WithKeys([]string{tag})
	res, err := c.Producer.SendSync(context.Background(), msg)
	if err != nil {
		klog.Errorf("Send message error: %s\n", err)
		return err
	}

	klog.V(5).Infof("Send message success: result=%s\n", res.String())
	return nil
}

func (c *RocketmqService) ProduceMsgWithTag(topic, tag string, data interface{}) error {
	msg := primitive.NewMessage(topic, Marshal(data))
	msg.WithTag(tag)
	res, err := c.Producer.SendSync(context.Background(), msg)
	if err != nil {
		klog.Errorf("Send message error: %s\n", err)
		return err
	}

	klog.V(5).Infof("Send message success: result=%s\n", res.String())
	return nil
}

func (c *RocketmqService) ConsumeMsgWithTag(topic string, tags []string) error {
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: strings.Join(tags, " || "),
	}
	err := c.Consumer.Subscribe(topic, selector, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		klog.V(5).Infof("subscribe callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		klog.Errorf(err.Error())
		return err
	}
	err = c.Consumer.Start()
	if err != nil {
		klog.Errorf("consumer start failed %v", err.Error())
		return err
	}
	return nil
}
func (c *RocketmqService) ConsumeMsg(topic string) error {
	selector := consumer.MessageSelector{}
	err := c.Consumer.Subscribe(topic, selector, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		klog.V(5).Infof("subscribe callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		klog.Errorf(err.Error())
	}
	err = c.Consumer.Start()
	if err != nil {
		klog.Errorf("start consumer error: %s", err.Error())
		return err
	}
	return nil
}

func (c *RocketmqService) ConsumeMsgDefine(topic string, f func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error {
	selector := consumer.MessageSelector{}
	err := c.Consumer.Subscribe(topic, selector, f)
	if err != nil {
		klog.Errorf(err.Error())
	}
	err = c.Consumer.Start()
	if err != nil {
		klog.Errorf("start consumer error: %s", err.Error())
		return err
	}
	return nil
}

func Marshal(data interface{}) []byte {
	marshal, err := json.Marshal(data)
	if err != nil {
		klog.Errorf("json marshal failed")
		return []byte{}
	}
	return marshal
}
