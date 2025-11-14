package repository

import (
	"github.com/lta2705/payment-processor/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(model *model.Transaction) error                                               // Create a new transaction record
	UpdateTransaction(model *model.Transaction) error                                               // Update an existing transaction record
	GetDB() *gorm.DB                                                                                // Get the underlying gorm DB instance
	FindByTransactionId(id *string) (*model.Transaction, error)                                     // Find a transaction by its ID
	FindByPcPosIdAndTransactionId(transactionId string, pcPosId string) (*model.Transaction, error) // Find a transaction by PcPosId and TransactionId
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

func (r *transactionRepositoryImpl) CreateTransaction(model *model.Transaction) error {
	return r.db.Create(model).Error
}
func (r *transactionRepositoryImpl) UpdateTransaction(model *model.Transaction) error {
	return r.db.Save(model).Error
}
func (r *transactionRepositoryImpl) GetDB() *gorm.DB {
	return r.db
}
func (r *transactionRepositoryImpl) FindByTransactionId(id *string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.Where("transaction_id = ?", id).First(&transaction).Error
	return &transaction, err
}
func (r *transactionRepositoryImpl) FindByPcPosIdAndTransactionId(transactionId string, pcPosId string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.Where("transaction_id = ? AND pc_pos_id = ?", transactionId, pcPosId).First(&transaction).Error
	return &transaction, err
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}