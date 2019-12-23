package routes

import (
	"fmt"
	"net/http"

	"../helpers"
)

// Sum Route to sum all values in the matrix
func Sum(responseWriter http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetMatrix(helpers.SUM, responseWriter, request)

	count := 0
	for n := range matrix {
		count += n
	}

	fmt.Fprint(responseWriter, count)
}
