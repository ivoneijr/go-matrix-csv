package routes

import (
	"fmt"
	"net/http"

	"../helpers"
)

// Sum Route to sum of the integers in the matrix
func Sum(w http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetRecords(w, request)
	flatted := helpers.MatrixToChannel(matrix)

	count := 0
	for n := range flatted {
		count += n
	}

	fmt.Fprint(w, count)
}
