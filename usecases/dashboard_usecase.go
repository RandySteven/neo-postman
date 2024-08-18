package usecases

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/models"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
)

type dashboardUsecase struct {
	dashboardRepo repositories_interfaces.DashboardRepository
}

func (d *dashboardUsecase) GetAvgResponseTimePerAPIs(ctx context.Context) (result []*models.AvgResponseTimePerApi, customErr *apperror.CustomError) {
	result, err := d.dashboardRepo.SelectAvgTimeResponseTime(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get avg response time`, err)
	}
	return result, nil
}

func (d *dashboardUsecase) GetExpectedUnexpectedResult(ctx context.Context) (result *models.ExpectedResultCount, customErr *apperror.CustomError) {
	result, err := d.dashboardRepo.SelectExpectedUnexpectedCount(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get dashboard result`, err)
	}
	return result, nil
}

var _ usecases_interfaces.DashboardUsecase = &dashboardUsecase{}

func NewDashboardUsecase(dashboardRepo repositories_interfaces.DashboardRepository) *dashboardUsecase {
	return &dashboardUsecase{
		dashboardRepo: dashboardRepo,
	}
}
