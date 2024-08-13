package handlers_interfaces

import "net/http"

type ApiCollectionHandler interface {
	UploadCollection(w http.ResponseWriter, r *http.Request)
	GetCollections(w http.ResponseWriter, r *http.Request)
	GetCollectionDetail(w http.ResponseWriter, r *http.Request)
}
