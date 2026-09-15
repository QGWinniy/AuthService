package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	http_user "AuthService/auth-service/internal/http/users"
	clientsessions "AuthService/auth-service/internal/infrastructure/sessions"
	clientusers "AuthService/auth-service/internal/infrastructure/users"

	"AuthService/auth-service/internal/usecase/sessions"
	"AuthService/auth-service/internal/usecase/users"
)

func main() {
	// =========================
	// User Service
	// =========================

	userConn, err := grpc.NewClient(
		"users-service:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer userConn.Close()

	usersClient := clientusers.NewClient(userConn)

	// =========================
	// Session Service
	// =========================

	sessionConn, err := grpc.NewClient(
		"sessions-service:50052",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sessionConn.Close()

	sessionClient := clientsessions.NewClient(sessionConn)

	// =========================
	// UseCases
	// =========================

	registerUC := usecase_users.NewRegisterUC(
		usecase_users.RegisterConfig{
			UserGateway: usersClient,
		},
	)

	loginUC := usecase_users.NewLoginUC(
		usecase_users.LoginConfig{
			UserGateway: usersClient,
		},
	)

	createSessionUC := usecase_sessions.NewCreateSessionUC(
		usecase_sessions.CreateConfig{
			SessionGateway: sessionClient,
		},
	)

	// =========================
	// HTTP Handler
	// =========================

	usersHandler := http_user.NewUsersHandler(
		http_user.Config{
			RegisterUC:      registerUC,
			LoginUC:         loginUC,
			CreateSessionUC: createSessionUC,
		},
	)

	// =========================
	// Router
	// =========================

	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /register",
		usersHandler.Register,
	)

	mux.HandleFunc(
		"POST /login",
		usersHandler.Login,
	)

	log.Println("auth-service started on :8080")

	if err := http.ListenAndServe(
		":8080",
		mux,
	); err != nil {
		log.Fatal(err)
	}
}