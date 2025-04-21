package persistence

import "time"

type UserModel struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}

// ชื่อ table
func (UserModel) TableName() string {
	return "users"
}
