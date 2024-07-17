package handlers_interfaces

import "net/http"

type DevHandler interface {
	GetListUrl(w http.ResponseWriter, r *http.Request)
	DummyTester(w http.ResponseWriter, r *http.Request)
	Hello(w http.ResponseWriter, r *http.Request)
}
