package repository

import (
	"time"
)

type UserModel struct {
	ID        string `gorm:"primaryKey"`
	Uid       string `gorm:"uniqueIndex"`
	Name      string
	Email     string    `gorm:"uniqueIndex"`
	AvatarUrl string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (UserModel) TableName() string {
	return "users"
}
