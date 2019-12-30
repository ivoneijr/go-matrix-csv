package routes

import (
	"testing"
)

func TestInvert(t *testing.T) {
	response, err := getResponseFor(handlerParams{Invert, "/invert", "POST"})
	if err != nil {
		t.Fatal(err)
	}

	expected := "1,4,7\n2,5,8\n3,6,9\n"

	if response != expected {
		t.Errorf("Invert handler returned wrong content got %v want %v",
			response, expected)
	}
}
