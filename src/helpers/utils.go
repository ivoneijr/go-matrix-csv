package helpers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

// OperationType enum for operation types
type OperationType int

// Operation types definition
const (
	SUM      OperationType = 0
	FLATTEN  OperationType = 1
	MULTIPLY OperationType = 2
	INVERT   OperationType = 3
	ECHO     OperationType = 4
)

// transpose return a reversed [][]string
func Transpose(slice [][]string) [][]string {
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

// matrixToChannel
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

// getRecords get and parse csv file (form-data) to [][]string
func getRecords(w http.ResponseWriter, request *http.Request) ([][]string, error) {
	file, _, err := request.FormFile("file")

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error:  %s, please send csv file in form-data 'file' key", err.Error())))
		return nil, err
	}

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, err
	}

	defer file.Close()
	return records, nil
}

// GetMatrix return (<-chan int, [][]string) depending on the kind(type) param
func GetMatrix(
	operation OperationType,
	responseWriter http.ResponseWriter,
	request *http.Request,
) (<-chan int, [][]string) {
	matrix, _ := getRecords(responseWriter, request)

	switch operation {

	case SUM, FLATTEN, MULTIPLY:
		return matrixToChannel(matrix), nil

	case INVERT:
		return nil, Transpose(matrix)

	default:
		return nil, matrix
	}
}
