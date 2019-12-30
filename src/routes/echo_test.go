package routes

import (
	"testing"
)

func TestEcho(t *testing.T) {
	response, err := getResponseFor(handlerParams{Echo, "/echo", "POST"})
	if err != nil {
		t.Fatal(err)
	}

	expected := "1,2,3\n4,5,6\n7,8,9\n"

	if response != expected {
		t.Errorf("Echo handler returned wrong content got %v want %v",
			response, expected)
	}
}
