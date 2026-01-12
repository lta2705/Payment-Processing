package service

import (
	"github.com/lta2705/payment-processor/internal/model"
	"github.com/lta2705/payment-processor/internal/repository"
)

type TransactionService interface {
	CreateTransaction(transaction *model.Transaction) error
	UpdateTransaction(transaction *model.Transaction) error
	FindTransactionById(id *string) (*model.Transaction, error)
}

type TransactionServiceImpl struct {
	txRepo repository.PaymentRepository
}

func (t *TransactionServiceImpl) CreateTransaction(transaction *model.Transaction) error {
	return t.txRepo.CreatePayment(transaction)
}

func (t *TransactionServiceImpl) UpdateTransaction(transaction *model.Transaction) error {
	return t.txRepo.UpdatePayment(transaction)
}

func (t *TransactionServiceImpl) FindTransactionById(id *string) (*model.Transaction, error) {
	return t.txRepo.FindByTransactionId(id)
}

func NewTransactionService(txRepo repository.PaymentRepository) TransactionService {
	return &TransactionServiceImpl{
		txRepo: txRepo,
	}
}
