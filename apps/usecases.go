package apps

import (
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/RandySteven/neo-postman/pkg/postgres"
	"github.com/RandySteven/neo-postman/usecases"
)

type Usecases struct {
	TestDataUsecase usecases_interfaces.TestDataUsecase
}

func NewUsecases(repo *postgres.Repositories) *Usecases {
	return &Usecases{
		TestDataUsecase: usecases.NewTestDataUsecase(repo.TestDataRepo),
	}
}
