package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafkago.Reader
}

func NewConsumer(reader *kafkago.Reader) *Consumer {
	return &Consumer{
		reader: reader,
	}
}

func (c *Consumer) Read(
	ctx context.Context,
) (kafkago.Message, error) {
	return c.reader.ReadMessage(ctx)
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}