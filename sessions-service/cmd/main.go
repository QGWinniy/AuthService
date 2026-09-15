package main

import (
	"log"
	"net"

	sessionspb "AuthService/contracts/sessions"

	// "AuthService/sessions-service/internal/infrastructure/postgres"
	"AuthService/sessions-service/internal/infrastructure/redis"
	grpc_sessions "AuthService/sessions-service/internal/transport/grpc"
	"AuthService/sessions-service/internal/usecase"

	"google.golang.org/grpc"
)

func main() {
	// =========================
	// Repository
	// =========================

	sessionRepo, err := redis.NewRedisSessionRepository()

	// sessionRepo, err := postgres.NewPostgresSessionRepository()
	if err != nil {
		log.Fatal(err)
	}
	defer sessionRepo.Close()

	// =========================
	// UseCases
	// =========================

	createSessionUC := usecase.NewCreateSessionUC(
		usecase.CreateSessionConfig{
			SessionRepo: sessionRepo,
		},
	)

	getSessionUC := usecase.NewGetSessionUC(
		usecase.GetSessionConfig{
			SessionRepo: sessionRepo,
		},
	)

	deleteSessionUC := usecase.NewDeleteSessionUC(
		usecase.DeleteSessionConfig{
			SessionRepo: sessionRepo,
		},
	)

	// =========================
	// gRPC Server
	// =========================

	grpcServer := grpc.NewServer()

	sessionServer := grpc_sessions.NewServer(
		grpc_sessions.ServerConfig{
			UcCreateSession: createSessionUC,
			UcGetSession:    getSessionUC,
			UcDeleteSession: deleteSessionUC,
		},
	)

	sessionspb.RegisterSessionsServiceServer(
		grpcServer,
		sessionServer,
	)

	// =========================
	// Listener
	// =========================

	listener, err := net.Listen(
		"tcp",
		":50052",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("sessions-service started on :50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}