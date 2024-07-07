package handlers_interfaces

import "net/http"

type JiraIssueHandler interface {
	CreateJiraTicket(w http.ResponseWriter, r *http.Request)
	GetAllJiraTickets(w http.ResponseWriter, r *http.Request)
	GetSpecificJiraTicket(w http.ResponseWriter, r *http.Request)
}
