package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

type MyConsumer struct{}

var (
	myConsumer   *MyConsumer
	consumerOnce sync.Once
)

func NewConsumer() *MyConsumer {
	consumerOnce.Do(func() {
		consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
		if err != nil {
			fmt.Printf("fail to start consumer, err: %v\n", err)
		}
		myConsumer = &MyConsumer{}
		myConsumer.ConsumeComment(consumer)
	})
	return myConsumer
}

func (*MyConsumer) ConsumeComment(consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions("comment")
	if err != nil {
		fmt.Printf("fail to get list of partition: err%v\n", err)
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("comment", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err: %v\n", partition, err)
			return
		}
		go func(sarama.PartitionConsumer) {
			defer pc.AsyncClose()
			for {
				for msg := range pc.Messages() {
					fmt.Printf("Partition: %d Offset: %d Key: %v Value: %v", msg.Partition, msg.Offset, msg.Key, msg.Value)
					var comment entity.Comment
					err := json.Unmarshal([]byte(msg.Value), &comment)
					if err != nil {
						fmt.Printf("failed to unmarshal message, err: %v\n", err)
					}
					err = repository.NewCommentDaoInstance().AddComment(&comment)
					if err != nil {
						fmt.Printf("failed to add comment, err: %v\n", err)
					}
				}
			}
		}(pc)
	}
}
