package handlers

import "testing"

func TestAllHandlers(t *testing.T) {

	suite := []struct {
		params   handlerParams
		expected string
		message  string
	}{
		{
			handlerParams{HealthCheckHandler, "/health-check", "GET"},
			`{"alive": true}`,
			"HealthCheckHandler returned unexpected content: got %v want %v",
		},
		{
			handlerParams{Echo, "/echo", "POST"},
			"1,2,3\n4,5,6\n7,8,9\n",
			"Echo handler returned wrong content got: %v want %v",
		},
		{
			handlerParams{Flatten, "/flatten", "POST"},
			"1,2,3,4,5,6,7,8,9",
			"Flatten handler returned wrong content got: %v want %v",
		},
		{
			handlerParams{Invert, "/invert", "POST"},
			"1,4,7\n2,5,8\n3,6,9\n",
			"Invert handler returned wrong content got: %v want %v",
		},
		{
			handlerParams{Multiply, "/multiply", "POST"},
			"362880",
			"Multiply handler returned wrong content got: %v want %v",
		},
		{
			handlerParams{Sum, "/sum", "POST"},
			"45",
			"Sum handler returned wrong content got: %v want %v",
		},
	}

	for _, test := range suite {
		response, err := getResponseFor(test.params)
		if err != nil {
			t.Fatal(err)
		}

		if response != test.expected {
			t.Errorf(test.message, response, test.expected)
		}
	}
}
