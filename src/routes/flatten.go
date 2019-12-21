package routes

import (
	"fmt"
	"net/http"

	"../helpers"
)

func test(matrix [][]string) <-chan string {
	c := make(chan string)

	for _, row := range matrix {
		for _, column := range row {
			fmt.Println(column)
			// c <- column
		}
	}

	fmt.Println("inserted")
	close(c)

	return c
}

// Flatten route to matrix as a 1 line string, with values separated by commas.
func Flatten(w http.ResponseWriter, r *http.Request) {
	records, _ := helpers.GetRecords(w, r)

	test(records)

	// for i := range flatted {
	// 	fmt.Printf("%d ", i)
	// }

	// var response string

	// for _, row := range records {
	// 	response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	// }
	// fmt.Fprint(w, response)

	response := "flat response"
	fmt.Fprint(w, response)
}
