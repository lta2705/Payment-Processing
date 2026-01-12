package handler

import (
	"encoding/json"
	"log"

	"github.com/lta2705/payment-processor/internal/model"
	"github.com/lta2705/payment-processor/internal/service"
)

type CardPaymentHandler interface {
	HandleCardPayment(payload []byte) error
}

type CardPaymentHandlerImpl struct {
	paymentService service.PaymentService
}

func (h *CardPaymentHandlerImpl) HandleCardPayment(payload []byte) error {
	log.Printf("Processing card payment: %s", string(payload))

	// Parse JSON payload
	var transaction model.Transaction
	if err := json.Unmarshal(payload, &transaction); err != nil {
		log.Printf("Failed to unmarshal card payment payload: %v", err)
		return err
	}

	// Set transaction type
	transaction.TransactionType = "CARD"

	// Process the payment
	return h.paymentService.ProcessCardTransaction(payload)
}

func NewCardPaymentHandler(paymentService service.PaymentService) CardPaymentHandler {
	return &CardPaymentHandlerImpl{
		paymentService: paymentService,
	}
}
