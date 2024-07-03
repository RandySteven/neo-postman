package apps

import (
	usecases_interfaces "go-api-test/interfaces/usecases"
	"go-api-test/pkg/postgres"
	"go-api-test/usecases"
)

type Usecases struct {
	TestDataUsecase usecases_interfaces.TestDataUsecase
}

func NewUsecases(repo *postgres.Repositories) *Usecases {
	return &Usecases{
		TestDataUsecase: usecases.NewTestDataUsecase(repo.TestDataRepo),
	}
}
