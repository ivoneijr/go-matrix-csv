package routes

import (
	"fmt"
	"net/http"
	"strings"

	"../helpers"
)

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

// Invert route to matrix as a string in matrix format where the columns and rows are inverted
func Invert(w http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetRecords(w, request)
	transposed := transpose(matrix)
	var response string

	for _, row := range transposed {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}
