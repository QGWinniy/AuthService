package http_user

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}

