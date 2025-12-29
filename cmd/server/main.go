package main

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/lta2705/payment-processor/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. Khởi tạo app từ wire
	application, err := app.InitializeApp()
	if err != nil {
		logger.Fatal("Failed to initialize:", err)
	}

	// 2. Tạo context có thể hủy khi nhận tín hiệu OS
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 3. Run server on Goroutine in order not to block main thread
	go application.Start(ctx)

	// Đợi tín hiệu tắt máy
	<-ctx.Done()

	// 4. Graceful shutdown server
	application.Stop()
	logger.Info("Application exited cleanly")
}
