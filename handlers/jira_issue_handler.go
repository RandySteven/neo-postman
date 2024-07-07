package handlers

import (
	"context"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/enums"
	handlers_interfaces "github.com/RandySteven/neo-postman/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/RandySteven/neo-postman/utils"
	"github.com/google/uuid"
	"net/http"
)

type JiraIssueHandler struct {
	usecase usecases_interfaces.JiraIssueUseCase
}

func (j *JiraIssueHandler) CreateJiraTicket(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.CreateJiraIssueRequest{}
		dataKey = `jira`
	)
	if err := utils.BindRequest(r, request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, customErr := j.usecase.CreateJiraTicket(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get records`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusCreated, `success create ticket`, &dataKey, result, nil)
}

func (j *JiraIssueHandler) GetAllJiraTickets(w http.ResponseWriter, r *http.Request) {
}

func (j *JiraIssueHandler) GetSpecificJiraTicket(w http.ResponseWriter, r *http.Request) {
}

var _ handlers_interfaces.JiraIssueHandler = &JiraIssueHandler{}

func NewJiraIssueHandler(usecase usecases_interfaces.JiraIssueUseCase) *JiraIssueHandler {
	return &JiraIssueHandler{
		usecase: usecase,
	}
}
