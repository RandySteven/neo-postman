package apps

import (
	"github.com/RandySteven/neo-postman/middlewares"
	"github.com/gorilla/mux"
)

func RegisterMiddleware(r *mux.Router) *mux.Router {
	r.Use(
		middlewares.CorsMiddleware,
		middlewares.LoggingMiddleware,
		middlewares.TimeoutMiddleware,
	)
	return r
}
