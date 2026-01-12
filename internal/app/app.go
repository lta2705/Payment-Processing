package app

import "github.com/lta2705/payment-processor/internal/functionality"

type App struct {
	Producer functionality.Producer
	Consumer functionality.Consumer
}

func NewApp(producer functionality.Producer, consumer functionality.Consumer) *App {
	return &App{
		Producer: producer,
		Consumer: consumer,
	}
}
