package handler

import "github.com/segmentio/kafka-go"

type PaymentProcessHandler interface {
	ProcessPayment(msg kafka.Message)
}

type PaymentProcessHandlerImpl struct{}

func (h *PaymentProcessHandlerImpl) ProcessPayment(msg kafka.Message) {

}

func NewPaymentProcessHandler() PaymentProcessHandler {
	return &PaymentProcessHandlerImpl{}
}
