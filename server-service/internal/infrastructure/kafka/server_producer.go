package server_producer

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"AuthService/server-service/internal/broker"
	"AuthService/server-service/internal/domain"
	kafkaclient "AuthService/server-service/pkg/kafka"
)

type ServerBroker struct {
	producer *kafkaclient.Producer
}

func NewServerBroker(
	producer *kafkaclient.Producer,
) *ServerBroker {
	return &ServerBroker{
		producer: producer,
	}
}

func (b *ServerBroker) PublishServerCreationRequested(
	ctx context.Context,
	userServer *domain.UserServer,
) error {

	data, err := json.Marshal(broker.ToServerCreationRequestedEvent(userServer))
	if err != nil {
		return err
	}

	log.Print(data, userServer, userServer.UserId)

	return b.producer.Write(
		ctx,
		[]byte(strconv.Itoa(userServer.UserId)),
		data,
	)
}
