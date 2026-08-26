package http_user

import (
	"encoding/json"
	"net/http"
	"reg/auth-service/internal/usecase/users"
	"reg/auth-service/internal/usecase/sessions"
)

type Config struct {
	RegisterUC *usecase_users.RegisterUC
	LoginUC *usecase_users.LoginUC

	CreateSessionUC *usecase_sessions.CreateSessionUC
	GetSession *usecase_sessions.GetSessionUC
}


type UsersHandler struct {
	config Config
}

func NewUsersHandler(conf Config) *UsersHandler {
	return &UsersHandler{
		config: conf,
	}
}

func (h *UsersHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	userID, err := h.config.RegisterUC.Register(
		r.Context(),
		req.Login,
		req.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(RegisterResponse{
		UserID:  userID,
		Message: "registered successfully",
	})
}

func (h *UsersHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	user, err := h.config.LoginUC.Login(
		r.Context(),
		req.Login,
		req.Password,
	)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	session, err := h.config.CreateSessionUC.Create(
		r.Context(),
		user.Id,
	)
	if err != nil {
		http.Error(w, "failed to create session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session.SessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpiresAt,
	})

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(LoginResponse{
		UserID:  user.Id,
		Message: "logged in successfully",
	}); err != nil {
		return
	}
}

