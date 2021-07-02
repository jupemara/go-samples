package domain

// TODO: add more domain knowledge
type User struct {
	Id   string
	Name string
}

func NewUser(id, name string) *User {
	return &User{
		Id:   id,
		Name: name,
	}
}
