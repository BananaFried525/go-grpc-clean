package main

import (
	"log"
	"net"

	grpcAdapter "go-grpc-clean/internal/adapter/grpc"
	persistenceAdapter "go-grpc-clean/internal/adapter/persistence"
	"go-grpc-clean/internal/infra/persistence"
	proto "go-grpc-clean/internal/pb"
	"go-grpc-clean/internal/pkg/config"
	"go-grpc-clean/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.Init()

	db := persistence.NewGormDB()
	repo := persistenceAdapter.NewUserGormRepo(db)
	service := usecase.NewUserUseCase(repo)

	grpcServer := grpcAdapter.NewGRPCServer(service)

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
