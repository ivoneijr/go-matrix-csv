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

func TestFlatten(t *testing.T) {
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

	req, err := http.NewRequest("POST", "/flatten", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	//ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Flatten)
	handler.ServeHTTP(rr, req)

	response := rr.Body.String()
	expected := "1,2,3,4,5,6,7,8,9"

	if response != expected {
		t.Errorf("Flatten handler returned wrong content got %v want %v",
			response, expected)
	}
}
