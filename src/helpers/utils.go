package helpers

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"sync"
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

// IntToBigInt return a new big.Int based on int param
func IntToBigInt(n int) *big.Int {
	nString := strconv.FormatInt(int64(n), 10)
	nBig, _ := new(big.Int).SetString(nString, 10)

	return nBig
}

// Transpose return a reversed matrix [][]string
func Transpose(matrix [][]string) [][]string {
	qtColumns, qtRows := len(matrix[0]), len(matrix)
	transposed := make([][]string, qtColumns)
	wg := sync.WaitGroup{}

	for i := range transposed {
		transposed[i] = make([]string, qtRows)
	}

	for i := 0; i < qtColumns; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for j := 0; j < qtRows; j++ {
				wg.Add(1)

				go func(i, j int) {
					defer wg.Done()

					transposed[i][j] = matrix[j][i]
				}(i, j)
			}
		}(i)
	}

	wg.Wait()

	return transposed
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
		w.Write([]byte(fmt.Sprintf("error  %s, please send csv file in form-data 'file' key", err.Error())))
		return nil, err
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, err
	}

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
