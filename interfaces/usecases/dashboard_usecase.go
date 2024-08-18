package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/models"
)

type DashboardUsecase interface {
	GetExpectedUnexpectedResult(ctx context.Context) (result *models.ExpectedResultCount, customErr *apperror.CustomError)
	GetAvgResponseTimePerAPIs(ctx context.Context) (result []*models.AvgResponseTimePerApi, customErr *apperror.CustomError)
}
