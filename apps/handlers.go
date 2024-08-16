package apps

import (
	"github.com/RandySteven/neo-postman/handlers"
	handlers_interfaces "github.com/RandySteven/neo-postman/interfaces/handlers"
	"github.com/RandySteven/neo-postman/pkg/elastics"
	"github.com/RandySteven/neo-postman/pkg/postgres"
	"github.com/RandySteven/neo-postman/pkg/redis"
)

type Handlers struct {
	DevHandler        handlers_interfaces.DevHandler
	TestDataHandler   handlers_interfaces.TestDataHandler
	JiraIssueHandler  handlers_interfaces.JiraIssueHandler
	TestRecordHandler handlers_interfaces.TestRecordHandler
	DashboardHandler  handlers_interfaces.DashboardHandler
}

func NewHandlers(repo *postgres.Repositories, cache *redis.RedisClient, documentary *elastics.ESClient) *Handlers {
	usecases := NewUsecases(repo, cache, documentary)
	return &Handlers{
		TestDataHandler:   handlers.NewTestDataHandler(usecases.TestDataUsecase),
		JiraIssueHandler:  handlers.NewJiraIssueHandler(usecases.JiraIssueUsecase),
		TestRecordHandler: handlers.NewTestRecordHandler(usecases.TestRecordUsecase),
		DevHandler:        handlers.NewDevHandler(),
		DashboardHandler:  handlers.NewDashboardHandler(usecases.DashboardUsecase),
	}
}
