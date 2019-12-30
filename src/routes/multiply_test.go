package routes

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	response, err := getResponseFor(handlerParams{Multiply, "/multiply", "POST"})
	if err != nil {
		t.Fatal(err)
	}

	expected := "362880"

	if response != expected {
		t.Errorf("Multiply handler returned wrong content got %v want %v",
			response, expected)
	}
}
