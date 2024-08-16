package apps

import (
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/RandySteven/neo-postman/pkg/elastics"
	"github.com/RandySteven/neo-postman/pkg/postgres"
	"github.com/RandySteven/neo-postman/pkg/redis"
	"github.com/RandySteven/neo-postman/usecases"
)

type Usecases struct {
	TestDataUsecase   usecases_interfaces.TestDataUsecase
	JiraIssueUsecase  usecases_interfaces.JiraIssueUseCase
	TestRecordUsecase usecases_interfaces.TestRecordUseCase
	DashboardUsecase  usecases_interfaces.DashboardUsecase
}

func NewUsecases(
	repo *postgres.Repositories,
	cache *redis.RedisClient,
	documentary *elastics.ESClient) *Usecases {
	return &Usecases{
		TestDataUsecase:   usecases.NewTestDataUsecase(repo.TestDataRepo, repo.TestRecordRepo, cache.TestDataCache, documentary.TestDataDocumentary),
		JiraIssueUsecase:  usecases.NewJiraIssueUsecase(repo.JiraIssueRepo),
		TestRecordUsecase: usecases.NewTestRecordUsecase(repo.TestDataRepo, repo.TestRecordRepo),
		DashboardUsecase:  usecases.NewDashboardUsecase(repo.DashboardRepo),
	}
}
