package main

type KafkaListenerConfig struct {
	host      string
	port      int32
	partition int32
	topic     string
}
