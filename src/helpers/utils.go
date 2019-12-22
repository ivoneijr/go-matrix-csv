package helpers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
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

func matrixToChannel(matrix [][]string) <-chan int {
	ch := make(chan int, 1)

	go func() {
		defer close(ch)

		for _, row := range matrix {
			for _, column := range row {
				number, _ := strconv.Atoi(column)
				ch <- number
			}
		}
	}()

	return ch
}

func getRecords(w http.ResponseWriter, request *http.Request) ([][]string, error) {
	file, _, err := request.FormFile("file")
	defer file.Close()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, err
	}

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, err
	}

	return records, nil
}

// GetMatrix return (<-chan int, [][]string) depending by kind param
func GetMatrix(
	kind string,
	responseWriter http.ResponseWriter,
	request *http.Request,
) (<-chan int, [][]string) {
	matrix, _ := getRecords(responseWriter, request)

	switch kind {

	case "SUM", "FLATTEN", "MULTIPLY":
		return matrixToChannel(matrix), nil

	case "INVERT":
		return nil, transpose(matrix)

	default:
		return nil, matrix
	}
}
