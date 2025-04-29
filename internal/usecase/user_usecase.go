package usecase

import (
	"context"
	repository "go-grpc-clean/internal/domain/repository"
	domain "go-grpc-clean/internal/domain/user"

	"github.com/google/uuid"
)

type UserUseCaseImpl struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) domain.IUserRepository {
	return &UserUseCaseImpl{repo: r}
}

func (s *UserUseCaseImpl) GetUser(id string) (*domain.User, error) {
	return &domain.User{
		ID:    "123",
		Name:  "Alice",
		Email: "alice@example.com",
	}, nil
}

func (s *UserUseCaseImpl) CreateUser(u *domain.DraftCreateUser) (*domain.CreateUser, error) {
	ctx := context.Background()

	// gen uid for user
	uid := uuid.New().String()

	user, error := s.repo.Create(ctx, &repository.UserModel{
		ID:        "",
		Uid:       uid,
		Name:      u.Name,
		Email:     u.Email,
		AvatarUrl: "",
	})

	if error != nil {
		return nil, error
	}

	return &domain.CreateUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Uid:       user.Uid,
		AvatarUrl: user.AvatarUrl,
	}, nil
}
