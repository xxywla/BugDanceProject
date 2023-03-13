package service

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

type MyProducer struct{}

var (
	myProducer *MyProducer
	client     sarama.SyncProducer
	clientOnce sync.Once
)

func NewProducer() *MyProducer {
	clientOnce.Do(func() {
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Partitioner = sarama.NewRandomPartitioner
		config.Producer.Return.Successes = true

		// connect to kafka
		var err error
		client, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
		if err != nil {
			fmt.Println("producer closed, err:", err)
			return
		}
		// defer client.Close()
		myProducer = &MyProducer{}
	})
	return myProducer
}

func (*MyProducer) ProduceComment(comment string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = "comment"
	msg.Value = sarama.StringEncoder(comment)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
