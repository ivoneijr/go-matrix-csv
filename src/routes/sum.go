package routes

import (
	"fmt"
	"net/http"
)

// Sum Route to sum of the integers in the matrix
func Sum(w http.ResponseWriter, r *http.Request) {
	// TODO: should implement
	response := "sum response"
	fmt.Fprint(w, response)

}
