package persistence

import "context"

type UserRepository interface {
	Create(ctx context.Context, u *UserModel) error
	FindByID(ctx context.Context, id string) (*UserModel, error)
	FindByEmail(ctx context.Context, email string) (*UserModel, error)
}
