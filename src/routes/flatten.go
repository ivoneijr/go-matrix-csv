package routes

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"../helpers"
)

// Flatten route to get matrix as a 1 line string, with values separated by commas.
func Flatten(responseWriter http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetMatrix(helpers.FLATTEN, responseWriter, request)
	var buffer bytes.Buffer

	for n := range matrix {
		buffer.WriteString(fmt.Sprintf("%d,", n))
	}

	response := strings.TrimSuffix(buffer.String(), ",")

	fmt.Fprint(responseWriter, response)
}
