package handlers

import (
	handlers_interfaces "github.com/RandySteven/neo-postman/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"net/http"
)

type JiraIssueHandler struct {
	usecase usecases_interfaces.JiraIssueUseCase
}

func (j *JiraIssueHandler) CreateJiraTicket(w http.ResponseWriter, r *http.Request) {
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
