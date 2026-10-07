package http

import (
	"AuthService/ssh-connect-service/internal/domain"
	transport "AuthService/ssh-connect-service/internal/service"
	"AuthService/ssh-connect-service/internal/usecase"
	consoletoken "AuthService/ssh-connect-service/pkg/consoleToken"
	"encoding/json"
	"net/http"
)

type Config struct {
	MakeToken *usecase.MakeTokenUC
	RepoUserServer transport.ServerServiceGateway
	ConsoleTokenManager *consoletoken.Manager
}


type Handler struct {
	config Config
}

func NewUsersHandler(conf Config) *Handler {
	return &Handler{
		config: conf,
	}
}

func (h *Handler) Connect(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req ConnectRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	serverID := req.UserServerID

	ctx := r.Context()

	jwt, _ := r.Cookie(domain.CookieSessionJWT)

	sessionID, token, err := h.config.MakeToken.MakeToken(ctx, serverID, jwt.Value)

	if err != nil {
		http.Error(w, "Make Token ERROR", http.StatusBadRequest)
	}

	response := ConnectResponse{
		SessionID: sessionID,
		Token:     token,
		WSURL:     "wss://ssh.example.com/ws", // ПОМЕНЯТЬ
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}



