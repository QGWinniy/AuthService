package consumer

import (
	"AuthService/server-service-consumer/internal/broker"
	"AuthService/server-service-consumer/internal/usecase"
	"context"
	"log"
)

type ConfigConsumer struct {
	Broker broker.ServerBrokerConsumer
	CreateUserServerUC *usecase.CreateUserServerUC
}

type Consumer struct {
	config ConfigConsumer
}

func NewConsumer(conf ConfigConsumer) *Consumer {
	consumer := &Consumer{
		config: conf,
	}
	return consumer
}

func (c *Consumer) CheckBroker(ctx context.Context) error {
	log.Println("[Consumer] started")

	for {
		log.Println("[Consumer] waiting for Kafka message...")

		userServer, err := c.config.Broker.PublishServerGetRequested(ctx)
		if err != nil {
			log.Printf("[Consumer] Kafka error: %v", err)

			if ctx.Err() != nil {
				log.Printf("[Consumer] context cancelled: %v", ctx.Err())
				return ctx.Err()
			}

			continue
		}

		log.Printf("[Consumer] message received: %+v", userServer)

		err = c.config.CreateUserServerUC.CreateUserServer(ctx, userServer)
		if err != nil {
			log.Printf("[Consumer] create user server error: %v", err)
			continue
		}

		log.Println("[Consumer] user server created successfully")
	}
}