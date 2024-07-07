package usecases

import (
	"context"
	"encoding/json"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	jira_client "github.com/RandySteven/neo-postman/pkg/jira"
	"github.com/andygrunwald/go-jira"
	"sync"
)

type jiraIssueUsecase struct {
	jiraRepository repositories_interfaces.JiraIssueRepository
	jiraApiAction  jira_client.JiraAction
}

func (j *jiraIssueUsecase) CreateJiraTicket(ctx context.Context, request *requests.CreateJiraIssueRequest) (result *responses.CreateJiraIssueResponse, customErr *apperror.CustomError) {
	issue := jira.Issue{
		Fields: &jira.IssueFields{
			Project: jira.Project{
				Key: request.Project.Key,
			},
			Type: jira.IssueType{
				Name: request.IssueType.Name.ToString(),
			},
			Summary: request.Summary,
		},
	}
	response, err := j.jiraApiAction.CreateIssue(ctx, &issue)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create jira`, err)
	}
	jiraIssue := &models.JiraIssue{}

	var (
		wg    sync.WaitGroup
		errCh = make(chan *apperror.CustomError)
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		err = json.Unmarshal(jiraIssue.Request, &issue)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to create jira`, err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		err = json.Unmarshal(jiraIssue.Response, &response.Response.Body)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to create jira`, err)
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()
	responseMap := make(map[string]interface{})

	err = json.Unmarshal(jiraIssue.Response, &responseMap)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create jira`, err)
	}

	jiraIssue.Link = responseMap["self"].(string)

	jiraIssue, err = j.jiraRepository.Save(ctx, jiraIssue)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create jira`, err)
	}

	select {
	case customErr = <-errCh:
		return nil, customErr
	}

	return
}

func (j *jiraIssueUsecase) GetAllJiraTickets(ctx context.Context) (result []*responses.JiraIssueListResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.JiraIssueUseCase = &jiraIssueUsecase{}

func NewJiraIssueUsecase(
	jiraRepository repositories_interfaces.JiraIssueRepository,
) *jiraIssueUsecase {
	jira, err := jira_client.NewJiraClient()
	if err != nil {
		return nil
	}
	return &jiraIssueUsecase{
		jiraRepository: jiraRepository,
		jiraApiAction:  jira,
	}
}
