package usecase

import (
	repository "go-grpc-clean/internal/domain/persistence"
	domain "go-grpc-clean/internal/domain/user"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (s *UserUseCase) GetUser(id int32) (*domain.User, error) {
	return nil, nil
}

func (s *UserUseCase) CreateUser(u *domain.User) error {
	return nil
}
