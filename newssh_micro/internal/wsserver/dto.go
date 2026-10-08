package wsserver

import (
	"ssh/internal/domain"
)

// WsConnectRequest matches the connection payload returned by ssh-connect-service.
type WsConnectRequest struct {
	SessionID string `json:"session_id"`
	Token     string `json:"token"`
	WSURL     string `json:"ws_url"`
}

type WsMassageToServer struct {
	Command []byte `json:"command"`
}

func MakeTerminalInput(massage *WsMassageToServer) *domain.TerminalInput {
	return &domain.TerminalInput{
		Command: massage.Command,
	}
}

type WsMassageToUser struct {
	Result []byte `json:"result"`
}

func MakeWsMassageToUser(output *domain.TerminalOutput) *WsMassageToUser {
	return &WsMassageToUser{
		Result: output.Result,
	}
}
