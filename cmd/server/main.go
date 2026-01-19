package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lta2705/payment-processor/internal/app"
)

func main() {
	paymentApp, err := app.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		log.Println("Starting payment processor consumer...")
		if err := paymentApp.Consumer.ReadTransaction(ctx); err != nil {
			errChan <- err
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log.Println("Payment processor started successfully. Press Ctrl+C to shutdown.")

	select {
	case err := <-errChan:
		log.Printf("Payment processor consumer encountered a fatal error: %v", err)
	case sig := <-sigChan:
		log.Printf("Received signal: %v. Initiating graceful shutdown...", sig)
	}

	cancel()

	log.Println("Payment processor stopped.")
}
