package transport

import (
	"context"
	"encoding/json"
	"net"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/lta2705/payment-processor/internal/dto"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type SessionManager interface {
	Add(conn net.Conn, terminalId string)
	Remove(key string)
	CloseAll()
}

type Handler struct {
	Sessions      SessionManager
	KafkaProducer *kafka.Writer
	KafkaConsumer *kafka.Reader
	DB            *gorm.DB
}

func NewHandler(
	sessions SessionManager,
	producer *kafka.Writer,
	consumer *kafka.Reader,
	db *gorm.DB,
) *Handler {
	return &Handler{
		Sessions:      sessions,
		KafkaProducer: producer,
		KafkaConsumer: consumer,
		DB:            db,
	}
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {
	addr := conn.RemoteAddr().String()

	defer func() {
		h.Sessions.Remove(addr)
		conn.Close()
		logger.Info("Session closed:", addr)
	}()

	buf := make([]byte, 4096)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := conn.Read(buf)
			if err != nil {
				return
			}

			h.processMessage(conn, addr, buf[:n])
		}
	}
}

func (h *Handler) processMessage(conn net.Conn, addr string, data []byte) {
	var msg dto.RegisterDTO

	if err := json.Unmarshal(data, &msg); err != nil {
		logger.Error("Invalid message format", err)
		return
	}

	switch msg.MsgType {
	case "0":
		h.Sessions.Add(conn, msg.TerminalId)
		logger.Info("Session opened:", msg.TerminalId)

	case "2":
		// TODO: push Kafka / xử lý transaction

	default:
		logger.Warn("Unknown MsgType:", msg.MsgType)
	}

	conn.Write([]byte("ACK\n"))
}

func (h *Handler) Close() {
	h.Sessions.CloseAll()
}
