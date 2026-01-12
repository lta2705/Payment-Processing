package middleware

import (
	"fmt"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/lta2705/payment-processor/internal/model"
	"github.com/lta2705/payment-processor/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(cfg *config.DBConfig) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&model.Transaction{}); err != nil {
		logger.Fatal("Failed to auto-migrate database schema", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("Failed to get sql.DB instance", err)
	}

	sqlDB.SetMaxOpenConns(cfg.DBMaxConns)
	sqlDB.SetMaxIdleConns(cfg.DBIdleConn)
	sqlDB.SetConnMaxLifetime(cfg.DBConnMaxLifetime)

	logger.Info("Database connection established successfully")
	return db
}
