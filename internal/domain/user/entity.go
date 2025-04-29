package domain

type User struct {
	ID        string
	Name      string
	Email     string
	Uid       string
	AvatarUrl string
	CreatedAt string
	UpdatedAt string
}

type DraftCreateUser struct {
	Name  string
	Email string
}

type CreateUser struct {
	ID        string
	Uid       string
	Name      string
	Email     string
	AvatarUrl string
}
