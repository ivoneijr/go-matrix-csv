package routes

import (
	"fmt"
	"net/http"
	"strings"

	"../helpers"
)

// Echo route to echo matrix as a string
func Echo(responseWriter http.ResponseWriter, request *http.Request) {
	_, matrix := helpers.GetMatrix(helpers.ECHO, responseWriter, request)
	var response string

	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(responseWriter, response)
}
