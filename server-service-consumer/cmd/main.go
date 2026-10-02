package main

import (
	"context"
	"log"

	"AuthService/server-service-consumer/internal/infrastructure/consumer"
	serverproducer "AuthService/server-service-consumer/internal/infrastructure/kafka"
	"AuthService/server-service-consumer/internal/infrastructure/postgres"
	"AuthService/server-service-consumer/internal/usecase"
	kafkaclient "AuthService/server-service-consumer/pkg/kafka"
)

func main() {
	kafkaConfig := kafkaclient.InitKafkaConfig()
	kafkaReader := kafkaclient.InitKafkaReader(kafkaConfig)

	consumerKafka := kafkaclient.NewConsumer(kafkaReader)

	defer func() {
		if err := consumerKafka.Close(); err != nil {
			log.Printf("close Kafka consumer: %v", err)
		}
	}()

	serverBroker := serverproducer.NewServerBroker(consumerKafka)

	userServersRepo, err := postgres.NewPostgresRepository()
	if err != nil {
		log.Fatalf("create postgres repository: %v", err)
	}
	defer func() {
		if err := userServersRepo.Close(); err != nil {
			log.Printf("close postgres repository: %v", err)
		}
	}()

	createUserServerUC := usecase.NewCreateUserServerUC(
		usecase.CreateUserServerConfig{
			UserServerRepository: userServersRepo,
		},
	)

	ctx := context.Background()

	userServerConsumer := consumer.NewConsumer(
		consumer.ConfigConsumer{
			Broker:             serverBroker,
			CreateUserServerUC: createUserServerUC,
		},
	)

	if err := userServerConsumer.CheckBroker(ctx); err != nil {
		log.Printf("consumer stopped with error: %v", err)
	}

	log.Printf("server-service-consumer START V2")
}