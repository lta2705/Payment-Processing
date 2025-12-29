package middleware

import (
	"github.com/lta2705/payment-processor/utils"
	"time"

	"github.com/lta2705/payment-processor/pkg/config"
	"github.com/segmentio/kafka-go"
)

func CreateKafkaProducer(kafkaCfg *config.KafkaProducerConfig) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(kafkaCfg.BootstrapServers...),
		Topic:                  kafkaCfg.ProducerTopic,
		RequiredAcks:           kafka.RequiredAcks(utils.ParseAcks(kafkaCfg.Acks)),
		BatchSize:              kafkaCfg.MaxInFlight,
		AllowAutoTopicCreation: true,
	}
}

func CreateKafkaConsumer(kafkaCfg *config.KafkaConsumerConfig) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:           kafkaCfg.BootstrapServers,
		GroupID:           kafkaCfg.ConsumerGroupID,
		Topic:             kafkaCfg.ConsumerTopic,
		HeartbeatInterval: time.Duration(kafkaCfg.HeartbeatInterval) * time.Millisecond,
		StartOffset:       utils.FilterOffset(kafkaCfg.AutoOffsetReset),
		CommitInterval:    utils.FilterEnableAutoCommit(kafkaCfg.EnableAutoCommit),
		IsolationLevel:    utils.FilterIsolationLevel(kafkaCfg.IsolationLevel),
	})
}
