package common

import (
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"log"
	"time"
	"encoding/json"
	"go/src/server/entity"
)

var (
	nsqProducer *nsq.Producer
)

func Producer() *nsq.Producer {
	nsqIp := viper.GetString("nsq.ip")
	nsqHost := viper.GetString("nsq.port")
	if nsqProducer == nil {
		p, err := nsq.NewProducer(nsqIp + ":" + nsqHost, nsq.NewConfig())
		if err != nil {
			log.Print("producer create error:", err.Error())
			return nil
		}
		nsqProducer = p
	}
	return nsqProducer
}

func Publish(topic string, message []byte) {
	p := Producer()
	if p == nil {
		return
	}
	if err := p.Publish(topic, message); err != nil{
		log.Print("producer publish message error:", err.Error())
	}
}

func TestPublist() {
	nickName := "你的名字"
	password := "123456"
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	user := entity.NsqEntity{nickName, password, currentTime}
	b, _ := json.Marshal(user)
	Publish("test1", b)
}


