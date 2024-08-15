package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/entities/models"
)

type DashboardRepository interface {
	SelectExpectedUnexpectedCount(ctx context.Context) (result *models.ExpectedResultCount, err error)
}
