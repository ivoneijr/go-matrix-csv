package routes

import (
	"fmt"
	"net/http"
)

// Multiply route for product of the integers in the matrix
func Multiply(w http.ResponseWriter, r *http.Request) {
	// TODO: should implement
	response := "multiply response"
	fmt.Fprint(w, response)
}
