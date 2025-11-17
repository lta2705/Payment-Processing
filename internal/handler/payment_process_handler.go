package handler

type PaymentProcessHandler interface {
	ProcessPayment()
}

type PaymentProcessHandlerImpl struct {}

func (h *PaymentProcessHandlerImpl) ProcessPayment() {
	
}

func NewPaymentProcessHandler() PaymentProcessHandler {
	return &PaymentProcessHandlerImpl{}
}