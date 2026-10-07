package main

import (
	httpserver "AuthService/ssh-connect-service/internal/http"
	serverserviceclient "AuthService/ssh-connect-service/internal/infrastructure/serverServiceClient"
	"AuthService/ssh-connect-service/internal/middleware"
	"AuthService/ssh-connect-service/internal/usecase"
	consoletoken "AuthService/ssh-connect-service/pkg/consoleToken"
	jwtmanager "AuthService/ssh-connect-service/pkg/jwt"
	"log"
	stdhttp "net/http"
	"os"
)

func main() {
	userJWTSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(userJWTSecret) == 0 {
		log.Fatal("JWT_SECRET is required")
	}

	consoleJWTSecret := []byte(os.Getenv("SSH_CONSOLE_JWT_SECRET"))
	if len(consoleJWTSecret) == 0 {
		log.Fatal("SSH_CONSOLE_JWT_SECRET is required")
	}

	encryptionKey := []byte(os.Getenv("SSH_CONSOLE_ENCRYPTION_KEY"))
	if len(encryptionKey) != 16 && len(encryptionKey) != 24 && len(encryptionKey) != 32 {
		log.Fatal("SSH_CONSOLE_ENCRYPTION_KEY must be 16, 24, or 32 bytes")
	}

	serverServiceURL := os.Getenv("SERVER_SERVICE_URL")
	if serverServiceURL == "" {
		serverServiceURL = "http://server-service:8088"
	}

	consoleTokenManager, err := consoletoken.NewManager(consoleJWTSecret, encryptionKey)
	if err != nil {
		log.Fatalf("create console token manager: %v", err)
	}

	serverClient := serverserviceclient.NewServerServiceClient(serverServiceURL, nil)

	makeToken := usecase.NewMakeTokenUC(usecase.MakeTokenConfig{
		RepoUserServer:      serverClient,
		ConsoleTokenManager: consoleTokenManager,
	})

	handler := httpserver.NewUsersHandler(httpserver.Config{
		MakeToken:           makeToken,
		RepoUserServer:      serverClient,
		ConsoleTokenManager: consoleTokenManager,
	})

	jwt := jwtmanager.NewManager(userJWTSecret)
	mux := stdhttp.NewServeMux()
	mux.HandleFunc("POST /connect", middleware.AuthJWTMiddleware(handler.Connect, jwt))

	port := os.Getenv("SSH_CONNECT_SERVICE_PORT")
	if port == "" {
		port = "8090"
	}

	log.Printf("ssh-connect-service is listening on :%s, server-service=%s", port, serverServiceURL)
	if err := stdhttp.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
