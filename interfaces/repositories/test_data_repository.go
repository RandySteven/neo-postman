package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/entities/models"
)

type TestDataRepository interface {
	Repositories[models.TestData]
	DeletedUnsavedTestData(ctx context.Context) (err error)
}
