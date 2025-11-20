package repository

import (
	"github.com/lta2705/payment-processor/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(model *model.Transaction) error                                                   // Create a new transaction record
	UpdatePayment(model *model.Transaction) error                                                   // Update an existing transaction record
	GetDB() *gorm.DB                                                                                // Get the underlying gorm DB instance
	FindByTransactionId(id *string) (*model.Transaction, error)                                     // Find a transaction by its ID
	FindByPcPosIdAndTransactionId(transactionId string, pcPosId string) (*model.Transaction, error) // Find a transaction by PcPosId and TransactionId
}

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func (r *PaymentRepositoryImpl) CreatePayment(model *model.Transaction) error {
	return r.db.Create(model).Error
}
func (r *PaymentRepositoryImpl) UpdatePayment(model *model.Transaction) error {
	return r.db.Save(model).Error
}
func (r *PaymentRepositoryImpl) GetDB() *gorm.DB {
	return r.db
}
func (r *PaymentRepositoryImpl) FindByTransactionId(id *string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.Where("transaction_id = ?", id).First(&transaction).Error
	return &transaction, err
}
func (r *PaymentRepositoryImpl) FindByPcPosIdAndTransactionId(transactionId string, pcPosId string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.Where("transaction_id = ? AND pc_pos_id = ?", transactionId, pcPosId).First(&transaction).Error
	return &transaction, err
}

func NewTransactionRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}
