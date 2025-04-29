package mocks

import (
	"context"
	"go-grpc-clean/internal/domain/repository"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (r *UserRepository) Create(ctx context.Context, u *repository.UserModel) (*repository.UserModel, error) {
	args := r.Called(ctx, u)
	return args.Get(0).(*repository.UserModel), args.Error(1)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*repository.UserModel, error) {
	args := r.Called(ctx, email)
	return args.Get(0).(*repository.UserModel), args.Error(1)
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*repository.UserModel, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*repository.UserModel), args.Error(1)
}
