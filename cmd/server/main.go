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
	// Initialize the application
	paymentApp, err := app.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Create context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start consumer worker on a separate goroutine
	go func() {
		log.Println("Starting payment processor consumer...")
		paymentApp.Consumer.ReadTransaction(ctx)
	}()

	// Setup graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	log.Println("Payment processor started successfully. Press Ctrl+C to shutdown.")

	// Wait for interrupt signal
	<-c
	log.Println("Shutting down payment processor...")

	// Cancel context to stop consumer
	cancel()

	log.Println("Payment processor stopped.")
}
