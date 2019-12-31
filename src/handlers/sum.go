package handlers

import (
	"fmt"
	"net/http"

	"../helpers"
)

// Sum Route to sum all values in the matrix
func Sum(responseWriter http.ResponseWriter, request *http.Request) {
	chMatrix, _ := helpers.GetMatrix(helpers.SUM, responseWriter, request)
	count := 0

	for n := range chMatrix {
		count += n
	}

	fmt.Fprint(responseWriter, count)
}
