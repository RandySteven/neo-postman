package usecases

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/RandySteven/neo-postman/pkg/jira"
)

type jiraIssueUsecase struct {
	jiraRepository repositories_interfaces.JiraIssueRepository
	jiraApiAction  jira.JiraAction
}

func (j *jiraIssueUsecase) CreateJiraTicket(ctx context.Context, request *requests.CreateJiraIssueRequest) (result *responses.CreateJiraIssueResponse, customErr *apperror.CustomError) {
	return
}

func (j *jiraIssueUsecase) GetAllJiraTickets(ctx context.Context) (result []*responses.JiraIssueListResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.JiraIssueUseCase = &jiraIssueUsecase{}

func NewJiraIssueUsecase(
	jiraRepository repositories_interfaces.JiraIssueRepository,
	jiraApiAction jira.JiraAction,
) *jiraIssueUsecase {
	return &jiraIssueUsecase{
		jiraRepository: jiraRepository,
		jiraApiAction:  jiraApiAction,
	}
}
