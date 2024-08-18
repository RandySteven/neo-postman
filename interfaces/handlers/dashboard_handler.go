package handlers_interfaces

import "net/http"

type DashboardHandler interface {
	GetExpectedUnexpectedResult(w http.ResponseWriter, r *http.Request)
	GetAvgResponseTimePerAPIs(w http.ResponseWriter, r *http.Request)
	GetCountMethod(w http.ResponseWriter, r *http.Request)
	GetActiveTools(w http.ResponseWriter, r *http.Request)
}
