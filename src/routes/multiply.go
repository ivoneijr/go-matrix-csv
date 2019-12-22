package routes

import (
	"fmt"
	"net/http"

	"../helpers"
)

// Multiply route for product of the integers in the matrix
func Multiply(w http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetRecords(w, request)
	flatted := helpers.MatrixToChannel(matrix)

	count := 1
	for n := range flatted {
		count *= n
	}

	fmt.Fprint(w, count)
}
