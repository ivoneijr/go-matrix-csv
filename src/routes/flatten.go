package routes

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"../helpers"
)

func makeFlat(matrix [][]string) <-chan string {
	ch := make(chan string, 1)

	go func() {
		defer close(ch)

		for _, row := range matrix {
			for _, column := range row {
				ch <- column
			}
		}
	}()

	return ch
}

// Flatten route to matrix as a 1 line string, with values separated by commas.
func Flatten(w http.ResponseWriter, r *http.Request) {
	matrix, _ := helpers.GetRecords(w, r)

	flatted := makeFlat(matrix)

	var buffer bytes.Buffer

	for n := range flatted {
		buffer.WriteString(fmt.Sprintf("%s,", n))
	}

	response := strings.TrimSuffix(buffer.String(), ",")

	fmt.Fprint(w, response)
}
