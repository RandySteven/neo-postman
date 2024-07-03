package apps

import (
	"go-api-test/handlers"
	handlers_interfaces "go-api-test/interfaces/handlers"
	"go-api-test/pkg/postgres"
)

type Handlers struct {
	TestDataHandler handlers_interfaces.TestDataHandler
}

func NewHandlers(repo *postgres.Repositories) *Handlers {
	usecases := NewUsecases(repo)
	return &Handlers{
		TestDataHandler: handlers.NewTestDataHandler(usecases.TestDataUsecase),
	}
}
