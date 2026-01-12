package functionality

import "github.com/lta2705/payment-processor/internal/worker"

type Producer interface {
	ProduceMessage(msg string) error
}

type ProducerImpl struct {
	producer worker.KafkaProducerWorker
}

func (ps *ProducerImpl) ProduceMessage(msg string) error {
	return ps.producer.SendMessage(msg)
}

func NewProduce(producer worker.KafkaProducerWorker) Producer {
	return &ProducerImpl{
		producer: producer,
	}
}
