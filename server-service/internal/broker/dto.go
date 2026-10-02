package broker

import (
	"AuthService/server-service/internal/domain"
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
		UserID:   strconv.Itoa(int(userServer.UserId)),
		Addr:     userServer.Addr,
		Name:     userServer.Name,
		UserName: userServer.UserName,
		Password: userServer.Password,
	}
}