//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
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

func InitializeApp() (*App, error) {
	wire.Build(
		ProvideListener,
		NewSessionManager,
		NewApp,
	)
	return &App{}, nil
}
