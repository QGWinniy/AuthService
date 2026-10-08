package wsserver

import (
	"log"

	// "net"
	"net/http"
	// "ssh/internal/domain"
	// "ssh/internal/Terminal"
	"ssh/internal/domain"
	newsendtoserver "ssh/internal/infrastructure/newSendToServer"
	newsendtouser "ssh/internal/infrastructure/newSendToUser"
	"ssh/internal/queue"
	"ssh/internal/transport"
	maketerminal "ssh/internal/usecase/makeTerminal"
	consoletoken "ssh/pkg/consoleToken"

	"github.com/gorilla/websocket"
)

// type WSServer interface { // хз нахуя интерефей списал с урока мб лучше сделать просто структуру
// 	Start() error
// }

type Config struct {
	// Mux *http.ServeMux
	// Srv *http.Server
	WsUpg              *websocket.Upgrader
	RepoUserServer     transport.ServerServiceGateway
	MakeTerminal       *maketerminal.UseCase
	ConsoleTokenManager *consoletoken.Manager
}

type WsSrv struct {
	config Config
}

func NewWsServer(conf Config) *WsSrv {
	return &WsSrv{
		config: conf,
	}
}

// func (ws *wsSrv) Start() error {

//     ws.config.Mux.HandleFunc("/", ws.indexHandler)
//     ws.config.Mux.HandleFunc("/terminal", ws.terminalHandler)
//     ws.config.Mux.Handle(
//         "/static/",
//         http.StripPrefix(
//             "/static/",
//             http.FileServer(http.Dir("internal/wsserver/static")),
//         ),
//     )
//     ws.config.Mux.HandleFunc("/ws", ws.wsHandler)

//     log.Println("server started on :8080")

//     return ws.config.Srv.ListenAndServe()
// }

// func (ws *wsSrv) indexHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(
// 		w,
// 		r,
// 		"internal/wsserver/static/index.html",
// 	)
// }

// func (ws *wsSrv) terminalHandler(w http.ResponseWriter, r *http.Request) {
//     http.ServeFile(
//         w,
//         r,
//         "internal/wsserver/static/terminal.html",
//     )
// }

// всё что выше к хуям убрать

func (ws *WsSrv) WsHandler(w http.ResponseWriter, r *http.Request) { // в идеале избавиться от ws
	log.Println("start wsHandler")

	if ws.config.ConsoleTokenManager == nil {
		http.Error(w, "console token manager is unavailable", http.StatusInternalServerError)
		return
	}

	request := WsConnectRequest{Token: r.URL.Query().Get("token")}
	if request.Token == "" {
		http.Error(w, "missing console token", http.StatusUnauthorized)
		return
	}

	claims, err := ws.config.ConsoleTokenManager.Parse(request.Token)
	if err != nil {
		log.Printf("invalid console token: %v", err)
		http.Error(w, "invalid console token", http.StatusUnauthorized)
		return
	}

	password, err := ws.config.ConsoleTokenManager.Password(claims)
	if err != nil {
		log.Printf("failed to decrypt console token password: %v", err)
		http.Error(w, "invalid console token", http.StatusUnauthorized)
		return
	}
	userServer := &domain.UserServer{
		ID:       claims.ServerID,
		Addr:     claims.Addr,
		UserName: claims.UserName,
		Password: password,
	}

	conn, err := ws.config.WsUpg.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
		return
	}
	log.Println("Upgrade ok")

	log.Println("userServer ok")
	log.Printf("SSH credentials: addr=%q username=%q password=%q", userServer.Addr, userServer.UserName, userServer.Password)

	terminal, err := ws.config.MakeTerminal.MakeTerminal(userServer)

	if err != nil {
		log.Printf("error ws: %v", err)
		return
	}
	log.Println("MakeTerminal ok")

	toServer := queue.New[*domain.TerminalInput](100)
	toUser := queue.New[*domain.TerminalOutput](100)

	connection := &Connection{
		Conn:     conn,
		Server:   userServer,
		Terminal: terminal,

		ToServer: toServer,
		ToUser:   toUser,
	}

	log.Println("Connection ok")
	consumer := newsendtoserver.NewConsumer(newsendtoserver.Config{
		Q:        connection.ToServer,
		Terminal: connection.Terminal,
	})
	log.Println("NewConsumer ok")

	producer := newsendtouser.NewProducer(newsendtouser.Config{
		Q:        connection.ToUser,
		Terminal: connection.Terminal,
	})
	log.Println("NewProducer ok")

	go ws.readFromClient(connection)
	go ws.readFromServer(connection)

	go consumer.SendToServer()
	go producer.SendToUser()

}

func (ws *WsSrv) readFromClient(connection *Connection) {
	for {
		msg := new(WsMassageToServer)

		if err := connection.Conn.ReadJSON(msg); err != nil {
			log.Printf("readFromClient ReadJSON error: %v", err)
			break
		}

		log.Printf("readFromClient: command=%q", msg.Command)

		input := MakeTerminalInput(msg)

		log.Printf("readFromClient: pushing command=%q", input.Command)

		err := connection.ToServer.Push(input)

		if err != nil {
			return
		}

		log.Println("readFromClient: pushed")
	}

	connection.Close()
}

func (ws *WsSrv) readFromServer(connection *Connection) {
	for {
		msg, ok := <-connection.ToUser.Pop()
		if !ok {
			break
		}

		if msg == nil {
			log.Println("readFromServer: FUCKING NIL MESSAGE")

		} else {
			log.Printf("readFromServer: %q", msg.Result)

			if err := connection.Conn.WriteJSON(MakeWsMassageToUser(msg)); err != nil {
				return
			}
		}
	}
}
