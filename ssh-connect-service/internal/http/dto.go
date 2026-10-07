package http

type ConnectRequest struct {
	UserServerID int `json:"userserverid"`
}

type ConnectResponse struct {
	SessionID string `json:"session_id"`
	Token     string `json:"token"`
	WSURL     string `json:"ws_url"`
}