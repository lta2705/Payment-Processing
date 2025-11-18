package middleware

import (
	"time"

	"github.com/lta2705/payment-processor/pkg/config"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var logger = CreateLogger()

func parseAcks(acks string) int {
	switch acks {
	case "all":
		return -1
	case "1":
		return 1
	case "0":
		return 0
	default:
		logger.Info("invalid acks value '%s', using '1' as default", zap.String("acks", acks))
		return 1
	}
}

func filterOffset(autoOffsetReset string) int64 {
	switch autoOffsetReset {
	case "earliest":
		return kafka.FirstOffset
	case "latest":
		return kafka.LastOffset
	 default:
		return kafka.LastOffset
	}
}

func filterEnableAutoCommit(enableAutoCommit bool) time.Duration {
	if enableAutoCommit {
		return 1 * time.Second
	}
	return 0
}

func filterIsolationLevel(isolationLevel string) kafka.IsolationLevel {
	switch isolationLevel {
	case "read_uncommitted":
		return kafka.ReadUncommitted
	case "read_committed":
		return kafka.ReadCommitted
	default:
		return kafka.ReadCommitted
	}
}

func CreateKafkaProducer(kafkaCfg *config.KafkaProducerConfig) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: kafkaCfg.BootstrapServers,
		Topic:  kafkaCfg.ProducerTopic,
		RequiredAcks:   parseAcks(kafkaCfg.Acks),
		BatchSize: kafkaCfg.MaxInFlight,
	})
}

func CreateKafkaConsumer(kafkaCfg *config.KafkaConsumerConfig) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   kafkaCfg.BootstrapServers,
		GroupID:   kafkaCfg.ConsumerGroupID,
		Topic:     kafkaCfg.ConsumerTopic,
		HeartbeatInterval: time.Duration(kafkaCfg.HeartbeatInterval) * time.Millisecond,
		StartOffset: filterOffset(kafkaCfg.AutoOffsetReset),
		CommitInterval: filterEnableAutoCommit(kafkaCfg.EnableAutoCommit),
		IsolationLevel: filterIsolationLevel(kafkaCfg.IsolationLevel),
	})
}