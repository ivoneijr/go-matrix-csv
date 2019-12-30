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
	"testing"
)

func getCSVPath() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", nil
	}

	splitted := strings.Split(dir, "/src/routes")[0]
	path := splitted + "/docs/sample/matrix.csv"

	return path, nil
}

func TestEcho(t *testing.T) {
	path, err := getCSVPath()
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/echo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	//ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Echo)
	handler.ServeHTTP(rr, req)

	response := rr.Body.String()
	expected := "1,2,3\n4,5,6\n7,8,9\n"

	if response != expected {
		t.Errorf("Echo handler returned wrong content got %v want %v",
			response, expected)
	}
}
