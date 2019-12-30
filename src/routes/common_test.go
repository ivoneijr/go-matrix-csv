package routes

import (
	"os"
	"strings"
)

// getCSVPath return a filepath from sample matrix.csv
func getCSVPath() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", nil
	}

	splitted := strings.Split(dir, "/src/routes")[0]
	path := splitted + "/docs/sample/matrix.csv"

	return path, nil
}
