package grpc

import (
	"context"
	domain "go-grpc-clean/internal/domain/user"
	proto "go-grpc-clean/internal/pb"
)

type GRPCServer struct {
	proto.UnimplementedUserServiceServer
	usecase domain.IUserRepository
}

func NewGRPCServer(u domain.IUserRepository) *GRPCServer {
	return &GRPCServer{usecase: u}
}

func (s *GRPCServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	u, err := s.usecase.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{User: &proto.User{Id: u.ID, Name: u.Name, Email: u.Email}}, nil
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	u := &domain.DraftCreateUser{
		Name:  req.User.Name,
		Email: req.User.Email,
	}

	user, err := s.usecase.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{User: &proto.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Uuid:      user.Uid,
		AvatarUrl: user.AvatarUrl,
	}}, nil
}
