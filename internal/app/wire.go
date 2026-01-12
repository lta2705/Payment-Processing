//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/lta2705/payment-processor/internal/functionality"
	"github.com/lta2705/payment-processor/internal/handler"
	"github.com/lta2705/payment-processor/internal/middleware"
	"github.com/lta2705/payment-processor/internal/repository"
	"github.com/lta2705/payment-processor/internal/service"
	"github.com/lta2705/payment-processor/internal/worker"
	"github.com/lta2705/payment-processor/pkg/config"
)

var repositorySet = wire.NewSet(
	repository.NewTransactionRepository,
)

var serviceSet = wire.NewSet(
	service.NewPaymentService,
)

var handlerSet = wire.NewSet(
	handler.NewCardPaymentHandler,
	handler.NewQrPaymentHandler,
)

var ProducerSet = wire.NewSet(
	config.LoadKafkaProducerConfig,
	middleware.CreateKafkaProducer,
	worker.NewProducerWorker,
	functionality.NewProduce,
)

var ConsumerSet = wire.NewSet(
	config.LoadKafkaConsumerConfig,
	middleware.CreateKafkaConsumer,
	worker.NewConsumerWorker,
	functionality.NewConsumer,
)

var databaseSet = wire.NewSet(
	config.LoadDBConfig,
	middleware.SetupDatabase,
)

func InitializeApp() (*App, error) {
	wire.Build(
		databaseSet,
		repositorySet,
		serviceSet,
		handlerSet,
		ProducerSet,
		ConsumerSet,
		NewApp,
	)
	return &App{}, nil
}
