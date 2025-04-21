package domain

type Repository interface {
	GetByID(id int32) (*User, error)
	Create(user *User) error
}
