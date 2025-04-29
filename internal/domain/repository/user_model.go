package repository

import "time"

type UserModel struct {
	ID        string
	Uid       string
	Name      string
	Email     string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}
