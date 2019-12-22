package routes

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"../helpers"
)

// Flatten route to matrix as a 1 line string, with values separated by commas.
func Flatten(w http.ResponseWriter, r *http.Request) {
	matrix, _ := helpers.GetRecords(w, r)
	flatted := helpers.MatrixToChannel(matrix)
	var buffer bytes.Buffer

	for n := range flatted {
		buffer.WriteString(fmt.Sprintf("%d,", n))
	}

	response := strings.TrimSuffix(buffer.String(), ",")

	fmt.Fprint(w, response)
}
