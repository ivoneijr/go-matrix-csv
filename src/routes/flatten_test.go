package routes

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	response, err := getResponseFor(handlerParams{Flatten, "/flatten", "POST"})
	if err != nil {
		t.Fatal(err)
	}

	expected := "1,2,3,4,5,6,7,8,9"

	if response != expected {
		t.Errorf("Flatten handler returned wrong content got %v want %v",
			response, expected)
	}
}
