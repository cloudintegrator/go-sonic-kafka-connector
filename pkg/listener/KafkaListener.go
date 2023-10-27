package main

import (
	"github.com/IBM/sarama"
	"log"
	"os"
)

type KafkaListener struct {
	config *KafkaListenerConfig
}

func (l *KafkaListener) Listen() string {
	confg := sarama.NewConfig()
	brokers := []string{"localhost:29092"}
	consumer, err := sarama.NewConsumer(brokers, confg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("********** Kafka Consumer started **********")
	pc, err := consumer.ConsumePartition("DATA.TOPIC", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}
	stop := make(chan os.Signal, 1)
	run := true
	for run {
		select {
		case msg := <-pc.Messages():
			log.Println(msg.Value)
		case err := <-pc.Errors():
			log.Fatal(err)
		case <-stop:
			run = false
		}
	}
	return ""
}

var EXPORT_KAFKA_LISTENER KafkaListener
