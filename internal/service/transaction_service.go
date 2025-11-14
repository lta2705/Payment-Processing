package service

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

type TransactionService interface {
	processCardTransaction()
	processQRTransaction()

}
type TransactionServiceImpl struct {}