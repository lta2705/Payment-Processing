package config

import (
	"log"

	"github.com/joho/godotenv"
)

type KafkaProducerConfig struct { 
        BootstrapServers  []string
        ProducerTopic     string  
        Acks              string  
        Retries           int     
        MaxInFlight       int
        EnableIdempotence bool
}

type KafkaConsumerConfig struct {
        BootstrapServers  []string
        ConsumerGroupID   string
        ConsumerTopic     string
        AutoOffsetReset   string
        EnableAutoCommit  bool
        IsolationLevel    string
        HeartbeatInterval int
}

func LoadKafkaProducerConfig() *KafkaProducerConfig {
        // Load .env file
        loadConfig()

        return &KafkaProducerConfig{
                BootstrapServers:  []string{getEnv("KAFKA_BROKER", "localhost:9092")},
                ProducerTopic:     getEnv("KAFKA_PRODUCER_TOPIC", "producer_topic"),
                Acks:              getEnv("KAFKA_PRODUCER_ACKS", "all"),
                Retries:           getEnvAsInt("KAFKA_PRODUCER_RETRIES", 5),
                MaxInFlight:       getEnvAsInt("KAFKA_PRODUCER_MAX_IN_FLIGHT", 5),
                EnableIdempotence: getEnvAsBool("KAFKA_PRODUCER_ENABLE_IDEMPOTENCE", true),
        }
}

func LoadKafkaConsumerConfig() *KafkaConsumerConfig {
        // Load .env file
        loadConfig()
        return &KafkaConsumerConfig{
                BootstrapServers:  []string{getEnv("KAFKA_BROKER", "localhost:9092")},
                ConsumerGroupID:   getEnv("KAFKA_CONSUMER_GROUP_ID", "consumer_group"),
                ConsumerTopic:     getEnv("KAFKA_CONSUMER_TOPIC", "consumer_topic"),
                AutoOffsetReset:   getEnv("KAFKA_CONSUMER_AUTO_OFFSET_RESET", "earliest"),
                EnableAutoCommit:  getEnvAsBool("KAFKA_CONSUMER_ENABLE_AUTO_COMMIT", true),
                IsolationLevel:    getEnv("KAFKA_CONSUMER_ISOLATION_LEVEL", "read_committed"),
                HeartbeatInterval: getEnvAsInt("KAFKA_CONSUMER_HEARTBEAT_INTERVAL_MS", 20000),
        }
}

func loadConfig() {
        err := godotenv.Load("C:\\Users\\Alliex\\Desktop\\Thesis\\Payment-Processing\\.env")
        if err != nil {
                log.Println(".env file not found, using environment variables")
        }

}