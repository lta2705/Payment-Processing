//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/lta2705/payment-processor/internal/middleware"
	"github.com/lta2705/payment-processor/internal/service"
	"github.com/lta2705/payment-processor/internal/transport"
	"github.com/lta2705/payment-processor/pkg/config"
	"net"
	"os"
)

func ProvideListener() (net.Listener, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}
	return net.Listen("tcp", ":"+port)
}

var DatabaseSet = wire.NewSet(
	config.LoadDBConfig,
	middleware.SetupDatabase,
)

var ProducerWorkerSet = wire.NewSet(
	config.LoadKafkaProducerConfig,
	middleware.CreateKafkaProducer,
)

var ConsumerWorkerSet = wire.NewSet(
	config.LoadKafkaConsumerConfig,
	middleware.CreateKafkaConsumer,
	service.NewConsumerWorker,
)

var sessionSet = wire.NewSet(
	NewSessionManager,
	wire.Bind(
		new(transport.SessionManager),
		new(*SessionManager),
	),
)

var ServerSet = wire.NewSet(
	transport.NewServer,
	transport.NewHandler,
)

func InitializeApp() (*App, error) {
	wire.Build(
		ProvideListener,
		DatabaseSet,
		ProducerWorkerSet,
		ConsumerWorkerSet,
		sessionSet,
		ServerSet,
		NewApp,
	)
	return &App{}, nil
}
