package routes

import (
	"fmt"
	"net/http"
)

// Invert route to matrix as a string in matrix format where the columns and rows are inverted
func Invert(w http.ResponseWriter, r *http.Request) {
	// TODO: should implement
	response := "invert response"
	fmt.Fprint(w, response)
}
