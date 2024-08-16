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
