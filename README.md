# Payment Processing System

Hệ thống xử lý thanh toán sử dụng Go, Kafka, và PostgreSQL.

## Cấu trúc Project

```
cmd/server/          # Entry point của ứng dụng
cmd/test-producer/   # Test producer để gửi message
internal/
├── app/            # Application setup và dependency injection 
├── dto/            # Data Transfer Objects
├── functionality/  # Business logic interfaces
├── handler/        # Message handlers
├── middleware/     # Database, Kafka, Logger middleware
├── model/          # Database models
├── repository/     # Data access layer
├── service/        # Business services
├── worker/         # Kafka workers
pkg/config/         # Configuration management
utils/              # Utility functions
```

## Yêu cầu hệ thống

- Go 1.24 hoặc mới hơn
- PostgreSQL 12+
- Apache Kafka 2.8+
- Docker & Docker Compose (tuỳ chọn)

## Chạy nhanh (Quick Start)

1. **Khởi động services:**
```bash
make dev-start
# hoặc
docker-compose up -d
```

2. **Build và chạy ứng dụng:**
```bash
make run-server
# hoặc
go run cmd/server/main.go
```

3. **Test producer (terminal khác):**
```bash
make run-test
# hoặc  
go run cmd/test-producer/main.go
```

4. **Dừng services:**
```bash
make dev-stop
```

## Commands hữu ích

```bash
# Build tất cả
make all

# Generate Wire dependencies
make generate-wire

# Xem logs Docker
make docker-logs

# Run tests
make test

# Build cho production
make build-linux
```

## Cấu hình

### Database (.env)
- `DB_HOST`: Địa chỉ PostgreSQL (mặc định: 127.0.0.1)
- `DB_PORT`: Port PostgreSQL (mặc định: 5433) 
- `DB_USER`: Username PostgreSQL (mặc định: postgres)
- `DB_PASSWORD`: Password PostgreSQL
- `DB_NAME`: Tên database (mặc định: payment_db)

### Kafka (.env)
- `KAFKA_BROKER`: Địa chỉ Kafka broker (mặc định: localhost:9092)
- `KAFKA_PRODUCER_TOPIC`: Topic để gửi message
- `KAFKA_CONSUMER_TOPIC`: Topic để nhận message
- `KAFKA_CONSUMER_GROUP_ID`: Consumer group ID

## API

Hệ thống xử lý 2 loại thanh toán:
- **CARD**: Thanh toán bằng thẻ
- **QR**: Thanh toán bằng QR code

Messages được gửi qua Kafka với key tương ứng và payload JSON.

### Ví dụ JSON Payload

```json
{
  "transaction_type": "CARD",
  "curr_cd": "USD",
  "tot_tr_amt": 100.50,
  "tip_amt": 10.00,
  "pc_pos_id": "POS001",
  "transaction_id": "TXN001",
  "msg_type": "PAYMENT",
  "status": "PENDING"
}
```

## Cấu trúc Database

### Transaction Table
- `id`: UUID primary key
- `transaction_type`: Loại giao dịch (CARD/QR)
- `curr_cd`: Mã tiền tệ
- `tot_tr_amt`: Tổng số tiền
- `tip_amt`: Tiền tip
- `pc_pos_id`: ID thiết bị POS
- `transaction_id`: ID giao dịch
- `status`: Trạng thái giao dịch
- `created_at`: Thời gian tạo

## Development

### Build
```bash
go build ./cmd/server
# hoặc
make build
```

### Test
```bash
go test ./...
# hoặc
make test
```

### Generate Wire Dependencies
```bash
go generate ./internal/app
# hoặc
make generate-wire
```

## Troubleshooting

### Database Connection Error
Đảm bảo PostgreSQL đang chạy:
```bash
docker-compose ps
# Nếu không chạy:
docker-compose up -d postgres
```

### Kafka Connection Error
Đảm bảo Kafka và Zookeeper đang chạy:
```bash
docker-compose ps
# Nếu không chạy:
docker-compose up -d zookeeper kafka
```

### Wire Generation Issues
```bash
go install github.com/google/wire/cmd/wire@latest
go generate ./internal/app
```

## Kiến trúc

Hệ thống sử dụng:
- **Wire**: Dependency injection
- **Kafka**: Message queue  
- **GORM**: ORM cho PostgreSQL
- **Zap**: Structured logging
- **Clean Architecture**: Tách biệt business logic và infrastructure