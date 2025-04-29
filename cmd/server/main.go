package main

import (
	"log"
	"net"

	grpcAdapter "go-grpc-clean/internal/adapter/grpc"
	database "go-grpc-clean/internal/infra/db"
	"go-grpc-clean/internal/infra/repository"
	proto "go-grpc-clean/internal/pb"
	"go-grpc-clean/internal/pkg/config"
	"go-grpc-clean/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.Init()

	dbConnection := database.NewGormDB()
	userRepository := repository.NewUserGormRepo(dbConnection)
	userUsercase := usecase.NewUserUseCase(userRepository)

	grpcServer := grpcAdapter.NewGRPCServer(userUsercase)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, grpcServer)
	reflection.Register(server)

	log.Println("gRPC server is running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
