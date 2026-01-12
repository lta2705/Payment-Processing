package main

import (
	"encoding/json"
	"log"

	"github.com/lta2705/payment-processor/internal/app"
	"github.com/lta2705/payment-processor/internal/dto"
)

func main() {
	// Initialize the application
	paymentApp, err := app.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Create sample card transaction
	cardTransaction := dto.TransactionDTO{
		TransactionType: "CARD",
		CurrCd:          "USD",
		TotTrAmt:        100.50,
		TipAmt:          10.00,
		PcPosId:         "POS001",
		TransactionId:   "TXN001",
		MsgType:         "PAYMENT",
		Status:          "PENDING",
	}

	// Marshal to JSON
	cardPayload, err := json.Marshal(cardTransaction)
	if err != nil {
		log.Fatalf("Failed to marshal card transaction: %v", err)
	}

	// Send card transaction
	log.Println("Sending card transaction...")
	if err := paymentApp.Producer.ProduceMessage(string(cardPayload)); err != nil {
		log.Printf("Failed to send card transaction: %v", err)
	} else {
		log.Println("Card transaction sent successfully!")
	}

	// Create sample QR transaction
	qrTransaction := dto.TransactionDTO{
		TransactionType: "QR",
		CurrCd:          "USD",
		TotTrAmt:        75.25,
		TipAmt:          5.00,
		PcPosId:         "POS002",
		TransactionId:   "TXN002",
		MsgType:         "PAYMENT",
		Status:          "PENDING",
	}

	// Marshal to JSON
	qrPayload, err := json.Marshal(qrTransaction)
	if err != nil {
		log.Fatalf("Failed to marshal QR transaction: %v", err)
	}

	// Send QR transaction
	log.Println("Sending QR transaction...")
	if err := paymentApp.Producer.ProduceMessage(string(qrPayload)); err != nil {
		log.Printf("Failed to send QR transaction: %v", err)
	} else {
		log.Println("QR transaction sent successfully!")
	}

	log.Println("Test producer completed!")
}
