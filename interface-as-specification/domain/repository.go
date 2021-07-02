package domain

type IRepository interface {
	FindById(string) (*User, error)
}
