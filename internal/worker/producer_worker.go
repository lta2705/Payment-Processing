package worker

import (
	"context"
	"encoding/json"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/lta2705/payment-processor/internal/dto"
	"github.com/segmentio/kafka-go"
)

type KafkaProducerWorker interface {
	SendMessage(value string) error
}

type KafkaProducerWorkerImpl struct {
	Writer *kafka.Writer
}

func (pw *KafkaProducerWorkerImpl) SendMessage(value string) error {
	logger.Info("Received message from terminal", value)

	// Try to parse the message to determine the key
	var transaction dto.TransactionDTO
	key := "UNKNOWN"
	if err := json.Unmarshal([]byte(value), &transaction); err == nil {
		key = transaction.TransactionType
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}

	err := pw.Writer.WriteMessages(context.Background(), msg)

	if err == nil {
		logger.Info("Message sent successfully to Kafka", "key", key, "value", value)
	} else {
		logger.Error("Failed to send message to Kafka", err)
	}

	return err
}

func NewProducerWorker(writer *kafka.Writer) KafkaProducerWorker {
	return &KafkaProducerWorkerImpl{
		Writer: writer,
	}
}
