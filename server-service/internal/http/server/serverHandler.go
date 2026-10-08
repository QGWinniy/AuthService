package http_server

import (
	"AuthService/server-service/internal/domain"
	"AuthService/server-service/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type Config struct {
	CreationRequestedUC *usecase.ServerCreationRequestedUC
	GetUserServersUC    *usecase.GetUserServersUC
	GetUserServerUC    *usecase.GetUserServerUC
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

	ctx := r.Context()

	userIDUint, ok := ctx.Value(domain.UserIDKeyContext).(uint64)

	userID := int64(userIDUint) // опасная хуйня переделать в будущем

	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	userServer := AddServerRequestToUserServer(&req)
	userServer.UserId = userID

	err := h.config.CreationRequestedUC.ServerCreationRequested(
		ctx,
		userServer,
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

func (h *ServerHandler) GetUserServers(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	userIDUint, ok := ctx.Value(domain.UserIDKeyContext).(uint64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userServers, err := h.config.GetUserServersUC.GetUserServers(
		ctx,
		int64(userIDUint),
	)
	if err != nil {
		log.Printf("get user servers error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := UserServersToGetUserServersResponse(userServers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}

func (h *ServerHandler) GetUserServer(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	userIDUint, ok := ctx.Value(domain.UserIDKeyContext).(uint64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var request struct {
		UserServerID int `json:"user_server_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	userServer, err := h.config.GetUserServerUC.GetUserServer(
		ctx,
		int(userIDUint),
		request.UserServerID,
	)
	if err != nil {
		log.Printf("get user server error: %v", err)
		http.Error(w, "user server not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userServer)
}
