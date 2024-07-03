package apps

import (
	"github.com/RandySteven/neo-postman/handlers"
	handlers_interfaces "github.com/RandySteven/neo-postman/interfaces/handlers"
	"github.com/RandySteven/neo-postman/pkg/postgres"
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
