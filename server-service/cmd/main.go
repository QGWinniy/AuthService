package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpserver "AuthService/server-service/internal/http/server"
	serverproducer "AuthService/server-service/internal/infrastructure/kafka"
	"AuthService/server-service/internal/infrastructure/postgres"
	"AuthService/server-service/internal/middleware"
	"AuthService/server-service/internal/usecase"
	"AuthService/server-service/pkg/jwt"
	kafkaclient "AuthService/server-service/pkg/kafka"
)

func main() {
	kafkaConfig := kafkaclient.InitKafkaConfig()
	kafkaWriter := kafkaclient.InitKafkaWriter(kafkaConfig)
	producer := kafkaclient.NewProducer(kafkaWriter)
	defer func() {
		if err := producer.Close(); err != nil {
			log.Printf("close Kafka producer: %v", err)
		}
	}()

	serverBroker := serverproducer.NewServerBroker(producer)
	creationRequestedUC := usecase.NewServerCreationRequestedUC(
		usecase.ServerCreationRequestedConfig{Broker: serverBroker},
	)

	userServersRepo, err := postgres.NewPostgresRepository()
	if err != nil {
		log.Fatalf("create postgres repository: %v", err)
	}
	defer func() {
		if err := userServersRepo.Close(); err != nil {
			log.Printf("close postgres repository: %v", err)
		}
	}()
	getUserServersUC := usecase.NewGetUserServersUC(
		usecase.GetUserServersConfig{Repository: userServersRepo},
	)

	managerJWT := jwt.NewManager([]byte(os.Getenv("JWT_SECRET")))

	handler := httpserver.NewSeverHandler(httpserver.Config{
		CreationRequestedUC: creationRequestedUC,
		GetUserServersUC:    getUserServersUC,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /servers", middleware.AuthJWTMiddleware(handler.AddServer, managerJWT))
	mux.HandleFunc("GET /servers", middleware.AuthJWTMiddleware(handler.GetUserServers, managerJWT))

	server := &http.Server{
		Addr:    ":8088",
		Handler: mux,
	}

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	log.Println("server-service_v7 started on :8088")

	go func() {
		log.Printf("server-service is listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("serve HTTP: %v", err)
		}
	}()

	<-shutdownSignal
	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		log.Printf("shutdown HTTP server: %v", err)
	}
}
