package routes

import (
	"fmt"
	"net/http"

	"../helpers"
)

func makeFlat(matrix [][]string) <-chan string {
	c := make(chan string, 9)

	for _, row := range matrix {
		for _, column := range row {
			c <- column
		}
	}
	close(c)

	return c
}

// Flatten route to matrix as a 1 line string, with values separated by commas.
func Flatten(w http.ResponseWriter, r *http.Request) {
	records, _ := helpers.GetRecords(w, r)
	flatted := makeFlat(records)

	var response string

	for n := range flatted {
		if len(flatted) == 0 {
			response += fmt.Sprintf("%s", n)
		} else {
			response += fmt.Sprintf("%s,", n)
		}
	}

	fmt.Fprint(w, response)
}
