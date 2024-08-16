package handlers_interfaces

import "net/http"

type DashboardHandler interface {
	GetExpectedUnexpectedResult(w http.ResponseWriter, r *http.Request)
}
