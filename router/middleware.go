package router

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs information about incoming requests.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		log.Printf("Request: %s %s", req.Method, req.URL.Path)

		// Call the next handler in the chain
		next(w, req)

		// Calculate and log the time taken
		duration := time.Since(startTime)
		log.Printf("Processed in %s", duration)
	}
}



