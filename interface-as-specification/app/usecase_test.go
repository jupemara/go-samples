package app_test

import (
	"testing"

	"github.com/jupemara/go-samples/interface-as-specification/app"
	"github.com/jupemara/go-samples/interface-as-specification/repository"
)

func TestUserFindUsecase(t *testing.T) {
	kvs := map[string]struct {
		Id      string
		IsError bool
	}{
		"exists user":     {"user001", false},
		"non-exists user": {"not-exists", true},
	}
	for k, v := range kvs {
		usecase := app.NewUsecase(
			repository.NewMemory(),
		)
		err := usecase.Execute(v.Id)
		if v.IsError && err == nil {
			t.Fatalf(`case: %s. this case shuld be occurr error.`, k)
		}
		if !v.IsError && err != nil {
			t.Fatalf(`case: %s. unexpected error occurred: %s`, k, err)
		}
	}
}
