package repository

import (
	"fmt"

	"github.com/jupemara/go-samples/interface-as-specification/domain"
)

type Memory struct {
	data map[string]*domain.User
}

func NewMemory() *Memory {
	data := map[string]*domain.User{
		"user001": domain.NewUser("user001", "John, Smith"),
		"user002": domain.NewUser("user001", "Tarou, Yamada"),
	}
	return &Memory{data: data}
}

func (r *Memory) FindById(id string) (*domain.User, error) {
	v, ok := r.data[id]
	if !ok {
		return nil, fmt.Errorf(`couldn't find "%s" user`, id)
	}
	return v, nil
}
