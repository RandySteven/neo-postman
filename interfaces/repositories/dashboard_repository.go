package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/entities/models"
)

type DashboardRepository interface {
	SelectExpectedUnexpectedCount(ctx context.Context) (result *models.ExpectedResultCount, err error)
	SelectAvgTimeResponseTime(ctx context.Context) (result []*models.AvgResponseTimePerApi, err error)
	SelectCountApiMethod(ctx context.Context) (result *models.CountApiMethod, err error)
}
