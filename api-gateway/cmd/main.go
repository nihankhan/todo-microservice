package main

import (
	"api-gateway/api"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	r := api.NewRouter()

	fmt.Println("REST server running...")
	go func() {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Fatalf("failed to start REST server: %v", err)
		}
	}()

	// Start gRPC server
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	log.Println("Starting gRPC server on port 50051...")

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
