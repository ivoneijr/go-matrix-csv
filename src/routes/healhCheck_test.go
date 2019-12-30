package routes

import (
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	response, err := getResponseFor(handlerParams{HealthCheckHandler, "/health-check", "GET"})
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"alive": true}`
	if response != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response, expected)
	}
}
