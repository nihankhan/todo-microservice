package server

import (
	"api-gateway/proto"
	"log"
	"net"
	"user/internal/handlers"
	"user/internal/repository"
	"user/internal/service"

	"google.golang.org/grpc"
)

func StartGrpcServer() {
	repo := repository.NewInMUserRepo()
	userService := service.NewUserService(repo)
	grpcHandler := handlers.NewGrpcHandler(userService)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, grpcHandler)

	log.Println("Server is running on port :50052")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
