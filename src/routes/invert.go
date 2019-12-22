package routes

import (
	"fmt"
	"net/http"
	"strings"

	"../helpers"
)

// Invert route to get matrix as a string in matrix format where the columns and rows are inverted
func Invert(responseWriter http.ResponseWriter, request *http.Request) {
	_, matrix := helpers.GetMatrix("INVERT", responseWriter, request)
	var response string

	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(responseWriter, response)
}
