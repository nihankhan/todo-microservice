package server

import (
	"log"
	"net"
	"time"
	"user-service/internal/handlers"
	"user-service/internal/repository"
	"user-service/internal/service"
	"user-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer() {
	jwtSecret := "todo-microservice-nullbyte"
	tokenExpiry := time.Hour * 24

	repo := repository.NewInMUserRepo()
	userService := service.NewUserService(repo, jwtSecret, tokenExpiry)
	grpcHandler := handlers.NewGrpcHandler(userService)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	proto.RegisterUserServiceServer(grpcServer, grpcHandler)

	log.Println("Server is running on port :50052")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
