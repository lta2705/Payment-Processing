package service


type TransactionService interface {
	processCardTransaction()
	processQRTransaction()

}
type TransactionServiceImpl struct {}