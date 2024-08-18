package responses

import "github.com/RandySteven/neo-postman/entities/models"

type DashboardResponse struct {
	TestResultCount        *models.ExpectedResultCount     `json:"test_result_count,omitempty"`
	AvgResponseTimePerApis []*models.AvgResponseTimePerApi `json:"avg_response_time_per_apis,omitempty"`
}
