package service

import (
	"github.com/lta2705/payment-processor/internal/repository"
	"go.uber.org/zap"
)

type PaymentService interface {
	processCardTransaction()
	processQRTransaction()
}
type PaymentServiceImpl struct {
	txRepo repository.PaymentRepository
	logger *zap.Logger
}

func (p PaymentServiceImpl) processCardTransaction() {
}

func (p PaymentServiceImpl) processQRTransaction() {}

func NewPaymentService(txRepo repository.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{txRepo: txRepo}
}
