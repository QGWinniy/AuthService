package main

import (
	"log"
	"net"

	userspb "AuthService/contracts/users"

	"AuthService/users-service/internal/infrastructure/postgres"
	// "AuthService/users-service/internal/repository"
	grpc_users "AuthService/users-service/internal/transport/grpc"
	"AuthService/users-service/internal/usecase"

	"google.golang.org/grpc"
)


func main() {
	// =========================
	// Repository
	// =========================

	userRepo, err := postgres.NewPostgresUserRepository()
	if err != nil {
		log.Fatal(err)
	}
	defer userRepo.Close()

	// =========================
	// UseCases
	// =========================

	createUserUC := usecase.NewCreateUC(
		usecase.CreteConfig{
			UsersRepo: userRepo,
		},
	)

	getUserUC := usecase.NewGetUserByLoginUC(
		usecase.GetUserByLoginConfig{
			UsersRepo: userRepo,
		},
	)

	deleteUserUC := usecase.NewDeleteUC(
		usecase.DeleteConfig{
			UsersRepo: userRepo,
		},
	)

	// =========================
	// gRPC Server
	// =========================

	grpcServer := grpc.NewServer()

	userServer := grpc_users.NewServer(
		grpc_users.Config{
			UcCreateUser: createUserUC,
			UcGetUserByLogin: getUserUC,
			UcDeleteUser: deleteUserUC,
		},
	)

	userspb.RegisterUserServiceServer(
		grpcServer,
		userServer,
	)

	// =========================
	// Listener
	// =========================

	listener, err := net.Listen(
		"tcp",
		":50051",
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("user-service started on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}