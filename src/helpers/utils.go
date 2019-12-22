package helpers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

// MatrixToChannel convert matrix in to int channel
func MatrixToChannel(matrix [][]string) <-chan int {
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

// GetRecords get matrix [][]string from csv
func GetRecords(w http.ResponseWriter, request *http.Request) ([][]string, error) {
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
