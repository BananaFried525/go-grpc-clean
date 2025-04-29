package repository

import (
	"context"

	domain "go-grpc-clean/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(ctx context.Context, u *domain.UserModel) (*domain.UserModel, error) {
	model := &UserModel{
		Uid:       u.ID,
		Name:      u.Name,
		Email:     u.Email,
		AvatarUrl: u.AvatarUrl,
	}

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return nil, err
	}

	return &domain.UserModel{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
	}, nil
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (*domain.UserModel, error) {
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

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
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
