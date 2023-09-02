package tool

import (
	"fmt"

	"poker/config"
	"time"

	"github.com/nsqio/go-nsq"
)

type producer struct {
	producer *nsq.Producer
}

var Producer producer

func init() {

	p, err := nsq.NewProducer(config.Base.NsqHost, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	Producer.producer = p

}
func (p *producer) Publish(topic string, message string) (err error) {
	//defer p.producer.Stop()
	if message == "" {
		fmt.Println("message is empty")
		return nil
	}
	if err = p.producer.Publish(topic, []byte(message)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 延迟消息
func (p *producer) DeferredPublish(topic string, delay time.Duration, message string) (err error) {
	if message == "" {
		fmt.Println("message is empty")
		return nil
	}
	if err = p.producer.DeferredPublish(topic, delay, []byte(message)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func PushMessage(topic string, val string) {
	producer := Producer
	producer.Publish(topic, val)
}
