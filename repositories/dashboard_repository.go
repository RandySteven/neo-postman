package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/neo-postman/entities/models"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	"github.com/RandySteven/neo-postman/queries"
	"log"
)

type dashboardRepository struct {
	db *sql.DB
}

func (d *dashboardRepository) SelectAvgTimeResponseTime(ctx context.Context) (result []*models.AvgResponseTimePerApi, err error) {
	result = []*models.AvgResponseTimePerApi{}
	rows, err := d.db.QueryContext(ctx, queries.GetAvgResponseTimePerAPIQuery.ToString())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		response := &models.AvgResponseTimePerApi{}
		err = rows.Scan(&response.Uri, &response.AvgTime)
		if err != nil {
			return nil, err
		}
		result = append(result, response)
	}
	return result, nil
}

func (d *dashboardRepository) SelectExpectedUnexpectedCount(ctx context.Context) (result *models.ExpectedResultCount, err error) {
	result = &models.ExpectedResultCount{}
	err = d.db.QueryRowContext(ctx, queries.GetExpectedAndUnexpectedDataQuery.ToString()).Scan(&result.Expected, &result.Unexpected)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}

var _ repositories_interfaces.DashboardRepository = &dashboardRepository{}

func NewDashboardRepository(db *sql.DB) *dashboardRepository {
	return &dashboardRepository{
		db: db,
	}
}
