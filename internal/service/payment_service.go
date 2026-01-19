package service

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/google/uuid"
	"github.com/lta2705/payment-processor/internal/constants"
	"github.com/lta2705/payment-processor/internal/model"
	"github.com/lta2705/payment-processor/internal/repository"
	"log"
)

type PaymentService interface {
	ProcessCardTransaction(payload []byte) error
	ProcessQRTransaction(payload []byte) error
}

type PaymentServiceImpl struct {
	txRepo repository.PaymentRepository
}

func (p *PaymentServiceImpl) ProcessCardTransaction(payload []byte) error {
	logger.Info("Processing card transaction...")

	jsonParsed, err := gabs.ParseJSON(payload)
	if err != nil {
		return err
	}

	// 1. Kiểm tra dữ liệu an toàn trước khi ép kiểu
	transID, ok1 := jsonParsed.S("transactionId").Data().(string)
	pcPosID, _ := jsonParsed.S("pcPosId").Data().(string) // Giả sử bạn có pcPosId trong JSON
	if !ok1 {
		return fmt.Errorf("missing transactionId in payload")
	}

	// 2. Tìm kiếm transaction hiện có
	existModel, err := p.txRepo.FindByTransactionIdAndPcPosId(transID, pcPosID)
	if err != nil {
		// Nếu không tìm thấy, bạn phải dừng lại hoặc tạo mới,
		// không được gán giá trị vào existModel (vì nó đang nil)
		logger.Error("Original transaction not found in DB: %v", err)
		return err
	}

	// 3. Mapping dữ liệu từ JSON vào model (nếu cần cập nhật thêm)
	if val, ok := jsonParsed.S("transactionType").Data().(string); ok {
		existModel.TransactionType = val
	}

	// 4. Xử lý logic trạng thái
	if existModel.TransactionType == "SALE" {
		existModel.Status = constants.TxStatusSuccess
	} else if existModel.TransactionType == "VOID" {
		existModel.Status = constants.TxStatusVoided
	}

	existModel.UpdatedBy = "TCP_SERVER"
	existModel.ErrorCode = constants.ErrCodeNoErr
	existModel.ErrorDetail = constants.ErrDetailCode0

	// 5. Cập nhật vào DB
	if err := p.txRepo.UpdatePayment(existModel); err != nil {
		logger.Errorf("Failed to update transaction: %v", err)
		return err
	}

	return nil
}

func (p *PaymentServiceImpl) ProcessQRTransaction(payload []byte) error {
	logger.Info("Processing QR transaction...", payload)

	jsonParsed, err := gabs.ParseJSON(payload)
	if err != nil {
		return err
	}

	if jsonParsed.S("responseCode").Data().(string) != "00" {
		logger.Info("QR transaction failed with response code:", jsonParsed.S("responseCode").Data().(string))
		return nil
	}

	var transaction model.Transaction
	// Generate UUID if not provided
	if transaction.ID == uuid.Nil {
		transaction.ID = uuid.New()
	}

	transaction.TransactionId = jsonParsed.S("transactionId").Data().(string)

	existModel, err := p.txRepo.FindByTransactionId(&transaction.TransactionId)
	if err == nil {
		logger.Info("Cannot find original transaction: %s", transaction.ID)
	}

	existModel.UpdatedBy = "TCP_SERVER"
	existModel.Status = "SUCCESS"
	existModel.ErrorCode = constants.ErrCodeNoErr
	existModel.ErrorDetail = constants.ErrDetailCode0

	logger.Info("Saving new card transaction: %s", transaction.ID)
	if err := p.txRepo.UpdatePayment(existModel); err != nil {
		log.Printf("Failed to save card transaction: %v", err)
		return err
	}

	log.Printf("QR transaction processed successfully: %s", transaction.ID)
	return nil
}

func NewPaymentService(txRepo repository.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{txRepo: txRepo}
}
