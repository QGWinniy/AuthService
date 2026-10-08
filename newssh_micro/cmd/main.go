package main

import (
	"log"
	"net/http"
	"os"

	"ssh/internal/middleware"
	maketerminal "ssh/internal/usecase/makeTerminal"
	"ssh/internal/wsserver"
	consoletoken "ssh/pkg/consoleToken"
	jwtmanager "ssh/pkg/jwt"

	"github.com/gorilla/websocket"
)


func main() {
	makeTerminalUseCase := maketerminal.NewUseCase()
	userJWTSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(userJWTSecret) == 0 {
		log.Fatal("JWT_SECRET is required")
	}
	userJWTManager := jwtmanager.NewManager(userJWTSecret)

	consoleJWTSecret := []byte(os.Getenv("SSH_CONSOLE_JWT_SECRET"))
	if len(consoleJWTSecret) == 0 {
		log.Fatal("SSH_CONSOLE_JWT_SECRET is required")
	}
	encryptionKey := []byte(os.Getenv("SSH_CONSOLE_ENCRYPTION_KEY"))
	if len(encryptionKey) != 16 && len(encryptionKey) != 24 && len(encryptionKey) != 32 {
		log.Fatal("SSH_CONSOLE_ENCRYPTION_KEY must be 16, 24, or 32 bytes")
	}
	consoleTokenManager, err := consoletoken.NewManager(consoleJWTSecret, encryptionKey)
	if err != nil {
		log.Fatalf("create console token manager: %v", err)
	}

	// userServerRepo, err := userserverpostgres.NewUserServerPostgres()
	server := wsserver.NewWsServer(wsserver.Config{
		// Mux: m,
		// Srv: &http.Server{
		// 	Addr: ":8080",
		// 	Handler: m,
		// },
		WsUpg:              &websocket.Upgrader{},
		// RepoUserServer: userServerRepo,
		MakeTerminal:        makeTerminalUseCase,
		ConsoleTokenManager: consoleTokenManager,
	})

	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /ws",
		middleware.AuthJWTMiddleware(server.WsHandler, userJWTManager),
	)

	log.Println("server started on :8080 v15")

	if err := http.ListenAndServe(
		":8089",
		mux,
	); err != nil {
		log.Fatal(err)
	}
}
