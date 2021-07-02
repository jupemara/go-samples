package app

import (
	"log"

	"github.com/jupemara/go-samples/interface-as-specification/domain"
)

type UserFindUsecase struct {
	repository domain.IRepository
}

func NewUsecase(repository domain.IRepository) *UserFindUsecase {
	return &UserFindUsecase{
		repository: repository,
	}
}

func (u *UserFindUsecase) Execute(id string) error {
	user, err := u.repository.FindById(id)
	if err != nil {
		return err
	}
	log.Printf(`given id %s user's name is "%s"`, user.Id, user.Name)
	return nil
}
