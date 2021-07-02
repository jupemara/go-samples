package main

import (
	"log"

	"github.com/jupemara/go-samples/interface-as-specification/app"
	"github.com/jupemara/go-samples/interface-as-specification/repository"
)

func main() {
	app := app.NewUsecase(
		repository.NewMemory(),
	)
	err := app.Execute("user001")
	if err != nil {
		log.Fatalf(`unexpected error: %s`, err)
	}
}
