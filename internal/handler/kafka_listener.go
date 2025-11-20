package handler

import "github.com/segmentio/kafka-go"

type KafkaListener interface {
	Start()
}

type KafkaListenerImpl struct {
	reader  *kafka.Reader
	handler func(msg kafka.Message) error
}

func (k *KafkaListenerImpl) Start() {
	go func() {
		for {
			
		}
	}
}
