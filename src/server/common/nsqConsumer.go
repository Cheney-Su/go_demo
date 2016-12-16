package common

import (
	"github.com/nsqio/go-nsq"
	"log"
	"github.com/spf13/viper"
	"fmt"
)

var (
	nsqConsumer *nsq.Consumer
)

func Consume() *nsq.Consumer {
	nsqIp := viper.GetString("nsq.ip")
	nsqHost := viper.GetString("nsq.port")
	if nsqConsumer == nil {
		c, err := nsq.NewConsumer("test1", "nsqChanel", nsq.NewConfig())
		if err != nil {
			log.Print("consume create error:", err.Error())
		}
		c.AddHandler(&Handler{})
		if err := c.ConnectToNSQD(nsqIp + ":" + nsqHost); err != nil {
			log.Print("consume connect error:", err.Error())
		}
		nsqConsumer = c
	}

	return nsqConsumer
}

func TestConsume() {
	Consume()
}

type Handler struct {

}

func (*Handler) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}


