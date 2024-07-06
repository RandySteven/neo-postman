package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
)

type JiraIssueUseCase interface {
	CreateJiraTicket(ctx context.Context, request *requests.CreateJiraIssueRequest) (result *responses.CreateJiraIssueResponse, customErr *apperror.CustomError)
	GetAllJiraTickets(ctx context.Context) (result []*responses.JiraIssueListResponse, customErr *apperror.CustomError)
}
