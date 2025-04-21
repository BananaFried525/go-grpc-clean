package persistence

import (
	"context"

	domain "go-grpc-clean/internal/domain/persistence"

	"gorm.io/gorm"
)

type userGormRepo struct {
	db *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) domain.UserRepository {
	return &userGormRepo{db: db}
}

func (r *userGormRepo) Create(ctx context.Context, u *domain.UserModel) error {
	model := UserModel{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *userGormRepo) FindByID(ctx context.Context, id string) (*domain.UserModel, error) {
	var model UserModel
	if err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &domain.UserModel{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
	}, nil
}

func (r *userGormRepo) FindByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
	var model UserModel
	if err := r.db.WithContext(ctx).First(&model, "email = ?", email).Error; err != nil {
		return nil, nil // treat as not found
	}
	return &domain.UserModel{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
	}, nil
}
