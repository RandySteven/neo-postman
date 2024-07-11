package middlewares

import (
	"context"
	"net/http"
	"time"
)

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		// Use the new context in the request
		r = r.WithContext(ctx)

		// Use a done channel to signal when the request handling is complete
		done := make(chan struct{})

		// Handle the request in a separate goroutine
		go func() {
			defer close(done)
			next.ServeHTTP(w, r)
		}()

		// Wait for either the request handling to complete or the timeout to expire
		select {
		case <-done:
			// Request handling completed within the timeout
			return
		case <-ctx.Done():
			// Timeout occurred before request handling completed
			if ctx.Err() == context.DeadlineExceeded {
				http.Error(w, "Request timed out", http.StatusRequestTimeout)
				return
			}
		}
	})
}
