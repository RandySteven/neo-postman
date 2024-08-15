package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/neo-postman/entities/models"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
)

type dashboardRepository struct {
	db *sql.DB
}

func (d *dashboardRepository) SelectExpectedUnexpectedCount(ctx context.Context) (result *models.ExpectedResultCount, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.DashboardRepository = &dashboardRepository{}

func NewDashboardRepository(db *sql.DB) *dashboardRepository {
	return &dashboardRepository{
		db: db,
	}
}
