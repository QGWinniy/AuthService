package server_consumer

import (
	"context"
	"encoding/json"


	"AuthService/server-service-consumer/internal/broker"
	"AuthService/server-service-consumer/internal/domain"
	kafkaclient "AuthService/server-service-consumer/pkg/kafka"
)

type ServerBroker struct {
	consumer *kafkaclient.Consumer
}

func NewServerBroker(
	consumer *kafkaclient.Consumer,
) *ServerBroker {
	return &ServerBroker{
		consumer: consumer,
	}
}

// func (b *ServerBroker) PublishServerCreationRequested(
// 	ctx context.Context,
// 	userServer *domain.UserServer,
// ) error {

// 	data, err := json.Marshal(broker.ToServerCreationRequestedEvent(userServer))
// 	if err != nil {
// 		return err
// 	}

// 	log.Print(data, userServer, userServer.UserId)

// 	return b.producer.Write(
// 		ctx,
// 		[]byte(strconv.Itoa(userServer.UserId)),
// 		data,
// 	)
// }

func (b *ServerBroker) PublishServerGetRequested(
	ctx context.Context,
) (*domain.UserServer, error) {
	msg, err := b.consumer.Read(ctx)

	if err != nil {
		return nil, err
	}

	var event broker.ServerCreationRequestedEvent

	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return nil, err
	}

	userServer, err := broker.ToUserServer(&event)

	if err != nil {
		return nil, err
	}

	return userServer, nil
}
