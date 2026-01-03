package service

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/segmentio/kafka-go"
)

type ConsumerWorker struct {
	Reader *kafka.Reader
}

func SendMessage(writer *kafka.Writer, key, value string) error {
	logger.Info("Message sent", key, value)
	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)
}

func (cw *ConsumerWorker) ConsumeMessage(reader *kafka.Reader, handler func(msg kafka.Message) error) {
	for {
		msg, err := reader.FetchMessage(context.Background())
		if err != nil {
			logger.Error("fetch error", err)
			continue
		}

		// Logic Processing
		if err := handler(msg); err != nil {
			logger.Error("handler error", err)
			continue
		}

		//commit after processing
		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			logger.Error("commit error", err)
		}
	}
}

func NewConsumerWorker(reader *kafka.Reader) *ConsumerWorker {
	return &ConsumerWorker{
		Reader: reader,
	}
}
