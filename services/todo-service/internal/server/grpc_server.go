package server

import (
	"log"
	"net"
	"todo/internal/handlers"
	"todo/internal/repository"
	"todo/internal/service"
	"todo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartTodoServer() {
	repo := repository.NewTodoRepository()
	todoService := service.NewTodoService(repo)
	grpcHandler := handlers.NewGrpcHandler(todoService)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	proto.RegisterTodoServiceServer(grpcServer, grpcHandler)

	log.Println("ToDo gRPC server is running on port :50053")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to server Todo Service: %v", err)
	}
}
