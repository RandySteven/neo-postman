package apps

import (
	"github.com/RandySteven/neo-postman/enums"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	HandlerFunc func(w http.ResponseWriter, r *http.Request)

	EndpointRouter struct {
		path    string
		handler HandlerFunc
		method  string
	}
)

func RegisterEndpointRouter(path, method string, handler HandlerFunc) *EndpointRouter {
	return &EndpointRouter{path: path, handler: handler, method: method}
}

func NewEndpointRouters(h *Handlers) map[enums.RouterPrefix][]EndpointRouter {
	endpointRouters := make(map[enums.RouterPrefix][]EndpointRouter)

	endpointRouters[enums.TestDataPrefix] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodPost, h.TestDataHandler.CreateTestAPI),
		*RegisterEndpointRouter("", http.MethodGet, h.TestDataHandler.GetAllRecords),
		*RegisterEndpointRouter("/search", http.MethodGet, h.TestDataHandler.SearchHistory),
		*RegisterEndpointRouter("/{id}", http.MethodGet, h.TestDataHandler.GetDetailRecord),
		*RegisterEndpointRouter("/{id}/saved", http.MethodGet, h.TestDataHandler.SaveRecord),
		*RegisterEndpointRouter("/{id}/unsaved", http.MethodGet, h.TestDataHandler.UnsavedRecord),
	}

	endpointRouters[enums.JiraIssuePrefix] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodPost, h.JiraIssueHandler.CreateJiraTicket),
		*RegisterEndpointRouter("", http.MethodGet, h.JiraIssueHandler.GetAllJiraTickets),
		*RegisterEndpointRouter("/{id}", http.MethodGet, h.JiraIssueHandler.GetSpecificJiraTicket),
	}

	endpointRouters[enums.TestRecordPrefix] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodPost, h.TestRecordHandler.CreateTestRecord),
		*RegisterEndpointRouter("", http.MethodGet, h.TestRecordHandler.GetAllTestRecords),
		*RegisterEndpointRouter("/{id}", http.MethodGet, h.TestRecordHandler.GetTestRecordDetail),
	}

	endpointRouters[enums.DevPrefix] = []EndpointRouter{
		*RegisterEndpointRouter("/listurl", http.MethodGet, h.DevHandler.GetListUrl),
	}

	endpointRouters[enums.DashboardPrefix] = []EndpointRouter{
		*RegisterEndpointRouter("/result-count", http.MethodGet, h.DashboardHandler.GetExpectedUnexpectedResult),
		*RegisterEndpointRouter("/avg-response-time", http.MethodGet, h.DashboardHandler.GetAvgResponseTimePerAPIs),
		*RegisterEndpointRouter("/count-method", http.MethodGet, h.DashboardHandler.GetCountMethod),
		*RegisterEndpointRouter("/active-services", http.MethodGet, h.DashboardHandler.GetActiveTools),
	}

	return endpointRouters
}

func (h *Handlers) InitRouter(r *mux.Router) {
	mapRouters := NewEndpointRouters(h)

	testDataRouter := r.PathPrefix(enums.TestDataPrefix.ToString()).Subrouter()
	for _, router := range mapRouters[enums.TestDataPrefix] {
		testDataRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.TestDataPrefix.ToString())
	}

	jiraIssueRouter := r.PathPrefix(enums.JiraIssuePrefix.ToString()).Subrouter()
	for _, router := range mapRouters[enums.JiraIssuePrefix] {
		jiraIssueRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.JiraIssuePrefix.ToString())
	}

	testRecordRouter := r.PathPrefix(enums.TestRecordPrefix.ToString()).Subrouter()
	for _, router := range mapRouters[enums.TestRecordPrefix] {
		testRecordRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.TestRecordPrefix.ToString())
	}

	devRouter := r.PathPrefix(enums.DevPrefix.ToString()).Subrouter()
	for _, router := range mapRouters[enums.DevPrefix] {
		devRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.DevPrefix.ToString())
	}

	dashboardRouter := r.PathPrefix(enums.DashboardPrefix.ToString()).Subrouter()
	for _, router := range mapRouters[enums.DashboardPrefix] {
		dashboardRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.DashboardPrefix.ToString())
	}
}

func (router *EndpointRouter) RouterLog(prefix string) {
	log.Printf("%12s | %4s/ \n", router.method, prefix+router.path)
}
