package routes

import "testing"

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
