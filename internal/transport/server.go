package transport

import (
	"context"
	"net"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type Server struct {
	Handler *Handler
}

func NewServer(
	sessions SessionManager,
	producer *kafka.Writer,
	consumer *kafka.Reader,
	db *gorm.DB,
) *Server {
	return &Server{
		Handler: NewHandler(sessions, producer, consumer, db),
	}
}

func (s *Server) HandleConnection(ctx context.Context, conn net.Conn) {
	s.Handler.Handle(ctx, conn)
}

func (s *Server) Close() {
	s.Handler.Close()
}
