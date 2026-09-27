package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafkago.Writer
}

func NewProducer(writer *kafkago.Writer) *Producer {
	return &Producer{
		writer: writer,
	}
}

func (p *Producer) Write(
	ctx context.Context,
	key []byte,
	value []byte,
) error {
	return p.writer.WriteMessages(
		ctx,
		kafkago.Message{
			Key:   key,
			Value: value,
		},
	)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}