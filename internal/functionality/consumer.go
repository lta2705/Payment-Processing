package functionality

import (
	"context"
	"fmt"
	"strings"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/lta2705/payment-processor/internal/handler"
	"github.com/lta2705/payment-processor/internal/worker"
	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	ReadTransaction(ctx context.Context) error
}

type ConsumerImpl struct {
	Consumer    worker.KafkaConsumerWorker
	CardHandler handler.CardPaymentHandler
	QrHandler   handler.QrPaymentHandler
}

func (cs *ConsumerImpl) ReadTransaction(ctx context.Context) error {
	logger.Info("Consumer started and waiting for messages...")

	err := cs.Consumer.ConsumeMessage(func(msg kafka.Message) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		key := string(msg.Key)
		logger.Infof("Message received! Key: %s", key)

		switch {
		case strings.Contains(key, "CARD"):
			return cs.CardHandler.HandleCardPayment(msg.Value)

		case strings.Contains(key, "QR"):
			return cs.QrHandler.HandleQrPayment(msg.Value)

		default:
			logger.Warnf("Unknown transaction type: %s", key)
			return nil
		}
	})

	if err != nil {
		return fmt.Errorf("kafka consume error: %w", err)
	}

	return nil
}

func NewConsumer(consumer worker.KafkaConsumerWorker, cardHandler handler.CardPaymentHandler, qrHandler handler.QrPaymentHandler) Consumer {
	return &ConsumerImpl{
		Consumer:    consumer,
		CardHandler: cardHandler,
		QrHandler:   qrHandler,
	}
}
