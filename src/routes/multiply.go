package routes

import (
	"fmt"
	"net/http"

	"../helpers"
)

// Multiply route to get a product of the integers in the matrix
func Multiply(responseWriter http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetMatrix("MULTIPLY", responseWriter, request)

	count := 1
	for n := range matrix {
		count *= n
	}

	fmt.Fprint(responseWriter, count)
}
