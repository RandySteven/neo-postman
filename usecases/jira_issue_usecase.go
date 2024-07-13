package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	jira_client "github.com/RandySteven/neo-postman/pkg/jira"
	"github.com/andygrunwald/go-jira"
	"log"
	"sync"
)

type jiraIssueUsecase struct {
	jiraRepository repositories_interfaces.JiraIssueRepository
	jiraApiAction  jira_client.JiraAction
}

func (j *jiraIssueUsecase) CreateJiraTicket(ctx context.Context, request *requests.CreateJiraIssueRequest) (result *responses.CreateJiraIssueResponse, customErr *apperror.CustomError) {
	if j == nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `GAK ADA CONNECT KE JIRANYA ANYING`, errors.New("NGENTOD"))
	}
	project, _, err := j.jiraApiAction.GetClient().Project.Get("KAN")
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get project`, err)
	}
	assigneeUser, _, _ := j.jiraApiAction.GetClient().User.Get("randysteven12@gmail.com")
	issue := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: assigneeUser,
			Reporter: assigneeUser,
			Project:  *project,
			Type: jira.IssueType{
				Name: request.IssueType.Name.ToString(),
			},
			Summary:     request.Summary,
			Description: request.Description,
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
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to marshal request issue`, err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		err = json.Unmarshal(jiraIssue.Response, &response.Response.Body)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to unmarshal response body`, err)
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
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to marshall response map`, err)
	}

	jiraIssue.Link = responseMap["self"].(string)

	jiraIssue, err = j.jiraRepository.Save(ctx, jiraIssue)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save jira`, err)
	}

	result = &responses.CreateJiraIssueResponse{
		ID:   jiraIssue.ID,
		Link: jiraIssue.Link,
	}

	select {
	case customErr = <-errCh:
		return nil, customErr
	default:
		return result, nil
	}
}

func (j *jiraIssueUsecase) GetAllJiraTickets(ctx context.Context) (result []*responses.JiraIssueListResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.JiraIssueUseCase = &jiraIssueUsecase{}

func NewJiraIssueUsecase(
	jiraRepository repositories_interfaces.JiraIssueRepository,
) *jiraIssueUsecase {
	jiraApiAction, err := jira_client.NewJiraClient()
	if err != nil {
		log.Println("error", err)
		return nil
	}
	return &jiraIssueUsecase{
		jiraRepository: jiraRepository,
		jiraApiAction:  jiraApiAction,
	}
}
