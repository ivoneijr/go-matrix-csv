package routes

import (
	"fmt"
	"net/http"
	"strings"

	"../helpers"
)

// Echo route to matrix as a string in matrix format
func Echo(w http.ResponseWriter, request *http.Request) {
	records, _ := helpers.GetRecords(w, request)
	var response string

	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}
