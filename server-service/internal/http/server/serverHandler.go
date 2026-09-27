package http_server

import (
	"log"
	"AuthService/server-service/internal/usecase"
	"encoding/json"
	"net/http"
)

type Config struct {
	CreationRequestedUC *usecase.ServerCreationRequestedUC
	// RegisterUC *usecase_users.RegisterUC
	// LoginUC *usecase_users.LoginUC

	// CreateSessionUC *usecase_sessions.CreateSessionUC
	// GetSession *usecase_sessions.GetSessionUC
}


type ServerHandler struct {
	config Config
}

func NewSeverHandler(conf Config) *ServerHandler {
	return &ServerHandler{
		config: conf,
	}
}

func (h *ServerHandler) AddServer(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req AddServerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	ctx := r.Context()


	err := h.config.CreationRequestedUC.ServerCreationRequested(
		ctx,
		AddServerRequestToUserServer(&req),
	)

	if err != nil {
		log.Printf("server creation error: %v", err)
	}

	response := AddServerResponse{
		ServerID: "stub-server-id",
		Message:  "server added successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(response)
}

// DeleteServer — заглушка удаления сервера.
func (h *ServerHandler) DeleteServer(
	w http.ResponseWriter,
	r *http.Request,
) {
	// TODO: реализовать удаление сервера

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(DeleteServerResponse{
		Message: "server deleted successfully",
	})
}

// GetUserServers — заглушка получения всех серверов пользователя.
func (h *ServerHandler) GetUserServers(
	w http.ResponseWriter,
	r *http.Request,
) {
	// TODO: реализовать получение серверов пользователя

	response := GetUserServersResponse{
		Servers: []ServerResponse{
			{
				ID:   "stub-server-id-1",
				Name: "Test Server",
			},
			{
				ID:   "stub-server-id-2",
				Name: "Another Server",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}

