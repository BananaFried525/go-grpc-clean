package usecase_test

import (
	"testing"

	"go-grpc-clean/internal/usecase"
	"go-grpc-clean/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_Success(t *testing.T) {
	repo := new(mocks.UserRepository)
	userUseCase := usecase.NewUserUseCase(repo)

	user, err := userUseCase.GetUser("123")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "123", user.ID)
}

func TestGetUser_Fail_EmptyID(t *testing.T) {
	repo := new(mocks.UserRepository)
	userUseCase := usecase.NewUserUseCase(repo)

	user, err := userUseCase.GetUser("")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "id_is_empty", err.Error())
}
