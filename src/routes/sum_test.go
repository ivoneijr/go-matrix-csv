package routes

import (
	"testing"
)

func TestSum(t *testing.T) {
	response, err := getResponseFor(handlerParams{Sum, "/sum", "POST"})
	if err != nil {
		t.Fatal(err)
	}

	expected := "45"

	if response != expected {
		t.Errorf("Multiply handler returned wrong content got %v want %v",
			response, expected)
	}
}
