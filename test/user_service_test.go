package usecase_test

import (
	"context"
	"testing"
	"time"

	"go-grpc-clean/internal/domain"
	"go-grpc-clean/internal/usecase"

	"github.com/stretchr/testify/assert"
)

// Fake Repository
type fakeRepo struct {
	users map[string]*domain.User
}

func (f *fakeRepo) Create(ctx context.Context, u *domain.User) error {
	f.users[u.ID] = u
	return nil
}
func (f *fakeRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {
	return f.users[id], nil
}
func (f *fakeRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	for _, u := range f.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, nil
}

func TestCreateUser_Success(t *testing.T) {
	repo := &fakeRepo{users: map[string]*domain.User{}}
	uc := usecase.NewUserUseCase(repo)

	user, err := uc.CreateUser(context.TODO(), "Alice", "alice@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
	assert.Equal(t, "alice@example.com", user.Email)
	assert.NotEmpty(t, user.ID)
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	repo := &fakeRepo{
		users: map[string]*domain.User{
			"1": {ID: "1", Email: "alice@example.com", Name: "Alice", CreatedAt: time.Now()},
		},
	}
	uc := usecase.NewUserUseCase(repo)

	_, err := uc.CreateUser(context.TODO(), "Bob", "alice@example.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email already exists")
}
