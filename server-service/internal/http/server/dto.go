package http_server

import "AuthService/server-service/internal/domain"

type AddServerRequest struct {
	Name     string `json:"name"`
	Addr     string `json:"addr"` 
	UserName string `json:"userName"`
	Password string `json:"password"`
	UserId   int    `json:"userId"`
}

func AddServerRequestToUserServer(r *AddServerRequest) *domain.UserServer {
	return &domain.UserServer{
		Addr:     r.Addr,
		Name:     r.Name,
		UserName: r.UserName,
		Password: r.Password,
		UserId:   r.UserId,
	}
}

type AddServerResponse struct {
	ServerID string `json:"server_id"`
	Message  string `json:"message"`
}

type DeleteServerResponse struct {
	Message string `json:"message"`
}

type ServerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetUserServersResponse struct {
	Servers []ServerResponse `json:"servers"`
}


