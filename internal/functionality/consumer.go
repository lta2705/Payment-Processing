package functionality

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/lta2705/payment-processor/internal/handler"
	"github.com/lta2705/payment-processor/internal/worker"
	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	ReadTransaction(ctx context.Context)
}

type ConsumerImpl struct {
	Consumer worker.KafkaConsumerWorker
	CardHandler  handler.CardPaymentHandler
    QrHandler handler.QrPaymentHandler
}

func (cs *ConsumerImpl) ReadTransaction(ctx context.Context) {
    logger.Info("Consumer started and waiting for messages...")

    err := cs.Consumer.ConsumeMessage(func(msg kafka.Message) error {
        // select {
        // case <-ctx.Done():
        //     return ctx.Err()
        // default:
        // }

        logger.Info("Message received!")

        key := string(msg.Key)
        
        switch key {
        case "CARD":
            return cs.CardHandler.HandleCardPayment(msg.Value)
        case "QR":
            return cs.QrHandler.HandleQrPayment(msg.Value)
        default:
            logger.Warnf("Received unknown message key: %s", key)
            return nil
        }
    })

    if err != nil {
        logger.Errorf("Consumer stopped with error: %v", err)
    }
}

func NewConsumer(consumer worker.KafkaConsumerWorker, cardHandler  handler.CardPaymentHandler, qrHandler handler.QrPaymentHandler) Consumer {
	return &ConsumerImpl{
		Consumer: consumer,
		CardHandler:  cardHandler,
		QrHandler: qrHandler,
	}
}