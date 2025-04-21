package grpc

import (
	"context"
	domain "go-grpc-clean/internal/domain/user"
	proto "go-grpc-clean/internal/pb"
	"go-grpc-clean/internal/usecase"
)

type GRPCServer struct {
	proto.UnimplementedUserServiceServer
	service *usecase.UserUseCase
}

func NewGRPCServer(s *usecase.UserUseCase) *GRPCServer {
	return &GRPCServer{service: s}
}

func (s *GRPCServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	u, err := s.service.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{User: &proto.User{Id: u.ID, Name: u.Name, Email: u.Email}}, nil
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	u := &domain.User{
		ID:    req.User.Id,
		Name:  req.User.Name,
		Email: req.User.Email,
	}
	err := s.service.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{User: req.User}, nil
}
