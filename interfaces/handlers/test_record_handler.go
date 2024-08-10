package handlers_interfaces

import "net/http"

type TestRecordHandler interface {
	CreateTestRecord(w http.ResponseWriter, r *http.Request)
	GetAllTestRecords(w http.ResponseWriter, r *http.Request)
	GetTestRecordDetail(w http.ResponseWriter, r *http.Request)
	SearchTestRecords(w http.ResponseWriter, r *http.Request)
}
