package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/entities/models"
)

type TestRecordRepository interface {
	Repositories[models.TestRecord]
	SaveSavedTestData(ctx context.Context) (err error)
}
