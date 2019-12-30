package routes

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestMultiply(t *testing.T) {
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

	req, err := http.NewRequest("POST", "/multiply", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	//ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Multiply)
	handler.ServeHTTP(rr, req)

	response := rr.Body.String()
	expected := "362880"

	if response != expected {
		t.Errorf("Multiply handler returned wrong content got %v want %v",
			response, expected)
	}
}
