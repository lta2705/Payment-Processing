package handler

import (
	"encoding/json"
	"log"

	"github.com/lta2705/payment-processor/internal/model"
	"github.com/lta2705/payment-processor/internal/service"
)

type QrPaymentHandler interface {
	HandleQrPayment(payload []byte) error
}

type QrPaymentHandlerImpl struct {
	paymentService service.PaymentService
}

func (q *QrPaymentHandlerImpl) HandleQrPayment(payload []byte) error {
	log.Printf("Processing QR payment: %s", string(payload))

	// Parse JSON payload
	var transaction model.Transaction
	if err := json.Unmarshal(payload, &transaction); err != nil {
		log.Printf("Failed to unmarshal QR payment payload: %v", err)
		return err
	}

	// Set transaction type
	transaction.TransactionType = "QR"

	// Process the payment
	return q.paymentService.ProcessQRTransaction(payload)
}

func NewQrPaymentHandler(paymentService service.PaymentService) QrPaymentHandler {
	return &QrPaymentHandlerImpl{
		paymentService: paymentService,
	}
}
