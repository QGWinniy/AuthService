package kafka

type Config struct {
	Brokers []string

	Topic string

	GroupID string

	MinBytes int
	MaxBytes int
}

