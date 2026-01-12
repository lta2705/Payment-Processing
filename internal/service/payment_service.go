package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lta2705/payment-processor/internal/model"
	"github.com/lta2705/payment-processor/internal/repository"
)

type PaymentService interface {
	ProcessCardTransaction(payload []byte) error
	ProcessQRTransaction(payload []byte) error
}

type PaymentServiceImpl struct {
	txRepo repository.PaymentRepository
}

func (p *PaymentServiceImpl) ProcessCardTransaction(payload []byte) error {
	log.Println("Processing card transaction...")

	// Parse transaction data
	var transaction model.Transaction
	if err := json.Unmarshal(payload, &transaction); err != nil {
		log.Printf("Failed to unmarshal transaction: %v", err)
		return err
	}

	// Generate UUID if not provided
	if transaction.ID == uuid.Nil {
		transaction.ID = uuid.New()
	}

	// Set defaults
	transaction.TransactionType = "CARD"
	transaction.Status = "PROCESSING"
	transaction.CreatedAt = time.Now()

	// Save to database
	if err := p.txRepo.CreatePayment(&transaction); err != nil {
		log.Printf("Failed to save card transaction: %v", err)
		return err
	}

	log.Printf("Card transaction processed successfully: %s", transaction.ID)
	return nil
}

func (p *PaymentServiceImpl) ProcessQRTransaction(payload []byte) error {
	log.Println("Processing QR transaction...")

	// Parse transaction data
	var transaction model.Transaction
	if err := json.Unmarshal(payload, &transaction); err != nil {
		log.Printf("Failed to unmarshal transaction: %v", err)
		return err
	}

	// Generate UUID if not provided
	if transaction.ID == uuid.Nil {
		transaction.ID = uuid.New()
	}

	// Set defaults
	transaction.TransactionType = "QR"
	transaction.Status = "PROCESSING"
	transaction.CreatedAt = time.Now()

	// Save to database
	if err := p.txRepo.CreatePayment(&transaction); err != nil {
		log.Printf("Failed to save QR transaction: %v", err)
		return err
	}

	log.Printf("QR transaction processed successfully: %s", transaction.ID)
	return nil
}

func NewPaymentService(txRepo repository.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{txRepo: txRepo}
}
