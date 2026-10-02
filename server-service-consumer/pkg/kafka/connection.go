package kafka

import (
	"os"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

func getEnvInt(
	key string,
	defaultValue int,
) int {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return result
}

func InitKafkaConfig() Config {
	return Config{
		Brokers: strings.Split(
			os.Getenv("KAFKA_BROKERS"),
			",",
		),
		Topic: os.Getenv("KAFKA_TOPIC"),
		GroupID: os.Getenv("KAFKA_GROUP_ID"),


		MinBytes: getEnvInt(
			"KAFKA_MIN_BYTES",
			10e3,
		),

		MaxBytes: getEnvInt(
			"KAFKA_MAX_BYTES",
			10e6,
		),
	}
}

func InitKafkaWriter(
	cfg Config,
) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(cfg.Brokers...),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func InitKafkaReader(
	cfg Config,
) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   cfg.Brokers,
		Topic:     cfg.Topic,
		GroupID:   cfg.GroupID,
		MinBytes:  cfg.MinBytes,
		MaxBytes:  cfg.MaxBytes,
	})
}

