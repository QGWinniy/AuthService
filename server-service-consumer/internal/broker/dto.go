package broker

import (
	"AuthService/server-service-consumer/internal/domain"
	"fmt"
	"strconv"
)

type ServerCreationRequestedEvent struct {
	UserID   string `json:"user_id"`
	Addr     string `json:"addr"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// ToServerCreationRequestedEvent конвертирует UserServer в ServerCreationRequestedEvent
func ToServerCreationRequestedEvent(userServer *domain.UserServer) ServerCreationRequestedEvent {
	return ServerCreationRequestedEvent{
		UserID:   strconv.Itoa(userServer.UserId),
		Addr:     userServer.Addr,
		Name:     userServer.Name,
		UserName: userServer.UserName,
		Password: userServer.Password,
	}
}

func ToUserServer(event *ServerCreationRequestedEvent) (*domain.UserServer, error) {
	userID, err := strconv.Atoi(event.UserID)
	if err != nil {
		return nil, fmt.Errorf("Error parse user_id")
	}

	return &domain.UserServer{
		UserId:   userID,
		Addr:     event.Addr,
		Name:     event.Name,
		UserName: event.UserName,
		Password: event.Password,
	}, nil
}

