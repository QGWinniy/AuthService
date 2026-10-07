package serverserviceclient

import "AuthService/ssh-connect-service/internal/domain"

type getUserServerByIDRequest struct {
	UserServerID int `json:"user_server_id"`
}

type userServerResponse struct {
	ID       int    `json:"id"`
	Addr     string `json:"addr"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (r userServerResponse) toDomain() *domain.UserServer {
	return &domain.UserServer{
		ID:       r.ID,
		Addr:     r.Addr,
		Name:     r.Name,
		UserName: r.UserName,
		Password: r.Password,
	}
}