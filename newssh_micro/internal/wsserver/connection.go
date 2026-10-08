package wsserver

import (
	terminal "ssh/internal/Terminal"
	"ssh/internal/domain"
	"ssh/internal/repository"

	"github.com/gorilla/websocket"
)

// type UserConnectServerWs struct {
// 	Conn *websocket.Conn
// 	UserServer domain.UserServer

// 	SSHStdin  io.WriteCloser
//     SSHStdout io.Reader

// 	SendToServer chan *wsMassageToServer
// 	SendToUser chan *wsMassageToUser
// }

type wsConnect struct {
	UserId int64 `json:"userid"`
	UserServerAddress string `json:"userserveraddress"`
}

type Connection struct {
	Conn *websocket.Conn 
	Server *domain.UserServer
	// userConnect *domain.UserStd

	ToServer repository.QueueMassage[*domain.TerminalInput]
	ToUser   repository.QueueMassage[*domain.TerminalOutput]

	Terminal terminal.Terminal
}

func (c *Connection) Close() { // по идеи это очень плохо зато очень удобно
	c.ToServer.Close()
	c.ToUser.Close()
	c.Terminal.Close()
}