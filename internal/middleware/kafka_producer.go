package middleware

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/lta2705/payment-processor/pkg/config"
	"go.uber.org/zap"
)

func CreateKafkaProducer(cfg *config.KafkaProducerConfig) (*kafka.Producer, error) {
	logger := CreateLogger()

	if cfg == nil {
		logger.Info("Producer Config is nil")
		return nil, nil
	}

	kafkaProdCfg := &kafka.ConfigMap{
		"bootstrap.servers":       cfg.BootstrapServers[0],
		"acks":                    cfg.Acks,
		"retries":                 cfg.Retries,
		"max.in.flight.requests":  cfg.MaxInFlight,
		"enable.idempotence":      cfg.EnableIdempotence,
	}

	producer, err := kafka.NewProducer(kafkaProdCfg)
	if err != nil {
		logger.Error("Failed to create Kafka producer", zap.Error(err))
		return nil, err
	}

	return producer, nil
}
