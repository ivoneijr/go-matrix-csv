package routes

import (
	"bytes"
	"fmt"
	"net/http"

	"../helpers"
)

func makeFlat(matrix [][]string) <-chan string {
	c := make(chan string, 9)
	defer close(c)

	for _, row := range matrix {
		// When I use it, I receive ERROR: "panic: send on closed channel"
		// go func(r []string, c chan string) {
		// 	for _, column := range r {
		// 		c <- column
		// 	}
		// }(row, c)
		for _, column := range row {
			c <- column
		}
	}

	return c
}

// Flatten route to matrix as a 1 line string, with values separated by commas.
func Flatten(w http.ResponseWriter, r *http.Request) {
	records, _ := helpers.GetRecords(w, r)
	flatted := makeFlat(records)

	var buffer bytes.Buffer

	for n := range flatted {
		if len(flatted) == 0 {
			buffer.WriteString(n)
		} else {
			buffer.WriteString(fmt.Sprintf("%s,", n))
		}
	}

	fmt.Fprint(w, buffer.String())
}
