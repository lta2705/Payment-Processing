package service

import (
	"context"

	"github.com/lta2705/payment-processor/internal/middleware"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var logger = middleware.CreateLogger()

func SendMessage(writer *kafka.Writer, key, value string) error {
	logger.Info("Message sent", zap.String("key", key), zap.String("value", value))
	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)
}

func ConsumeMessage(reader *kafka.Reader, handler func(msg kafka.Message) error) {
    for {
        msg, err := reader.FetchMessage(context.Background())
        if err != nil {
            logger.Error("fetch error", zap.Error(err))
            continue
        }

        // Xử lý logic
        if err := handler(msg); err != nil {
            logger.Error("handler error", zap.Error(err))
            continue
        }

        // Commit thủ công khi xử lý thành công
        if err := reader.CommitMessages(context.Background(), msg); err != nil {
            logger.Error("commit error", zap.Error(err))
        }
    }
}