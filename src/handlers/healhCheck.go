package handlers

import (
	"io"
	"net/http"
)

// HealthCheckHandler handler to /health-check route
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
