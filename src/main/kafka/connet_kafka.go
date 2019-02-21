package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

func trigger1() {
	brokerList := "localhost:9092" //定义kafka服务端地址
	groupID := "test-cons-1"  //定义消费者组
	topicList := "test"  //定义主题
	config := cluster.NewConfig()

	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // Does not work unless u use a different grp id

	consumer, err := cluster.NewConsumer(strings.Split(brokerList, ","), groupID, strings.Split(topicList, ","), config)
	if err != nil {
        logger.Printf("Failed to start consumer: %s\n", err)
	}

	go func() {
		for err := range consumer.Errors() {
			logger.Printf("Error: %s\n", err.Error())
		}
	}()

	go func() {
		for note := range consumer.Notifications() {
			logger.Printf("Rebalanced: %+v\n", note)
		}
	}()

	go func() {
		for msg := range consumer.Messages() {
			fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			consumer.MarkOffset(msg, "")
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-wait
	if err := consumer.Close(); err != nil {
		logger.Println("Failed to close consumer: ", err)
	}
}

func trigger2() {
	config := sarama.NewConfig()
	config.Consumer.Fetch.Max = 1

	cons, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		logger.Println("Error creating consumer ", err)
	}
	defer func() {
		if err := cons.Close(); err != nil {
			logger.Fatalln(err)
		}
	}()

	pc, err := cons.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		logger.Println(err)
	}
	defer func() {
		if err := pc.Close(); err != nil {
			logger.Fatalln(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-pc.Messages():
			{
				logger.Printf("Message : %s\nOffset : %d, Highwatermark: %d", string(msg.Value), msg.Offset, pc.HighWaterMarkOffset())
				consumed++
			}
		case <-signals:
			{
				break ConsumerLoop
			}
		}
	}
	<-signals
}

func main() {
	trigger2()
}