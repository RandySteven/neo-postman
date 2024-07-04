package handlers_interfaces

import "net/http"

type TestDataHandler interface {
	CreateTestAPI(w http.ResponseWriter, r *http.Request)
	GetAllRecords(w http.ResponseWriter, r *http.Request)
	GetDetailRecord(w http.ResponseWriter, r *http.Request)
}
