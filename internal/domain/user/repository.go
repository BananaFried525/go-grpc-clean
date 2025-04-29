package domain

type IUserRepository interface {
	GetUser(id string) (*User, error)
	CreateUser(user *DraftCreateUser) (*CreateUser, error)
}
