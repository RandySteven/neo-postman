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
}

func (router *EndpointRouter) RouterLog(prefix string) {
	log.Printf("%12s | %4s/ \n", router.method, prefix+router.path)
}
