package routes

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
)

// handlerParams struct definition
type handlerParams struct {
	handler func(http.ResponseWriter, *http.Request)
	route   string
	method  string
}

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

// gerResponseFor
func getResponseFor(params handlerParams) (string, error) {
	path, err := getCSVPath()
	if err != nil {
		return "", err
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(params.method, params.route, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	//ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(params.handler)
	handler.ServeHTTP(rr, req)

	return rr.Body.String(), nil
}
