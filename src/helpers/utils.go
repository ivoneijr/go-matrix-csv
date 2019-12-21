package helpers

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

// GetRecords get matrix from csv
func GetRecords(w http.ResponseWriter, r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
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
