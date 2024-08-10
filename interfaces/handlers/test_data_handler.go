package handlers_interfaces

import "net/http"

type TestDataHandler interface {
	CreateTestAPI(w http.ResponseWriter, r *http.Request)
	GetAllRecords(w http.ResponseWriter, r *http.Request)
	GetDetailRecord(w http.ResponseWriter, r *http.Request)
	SaveRecord(w http.ResponseWriter, r *http.Request)
	UnsavedRecord(w http.ResponseWriter, r *http.Request)
	SearchHistory(w http.ResponseWriter, r *http.Request)
}
