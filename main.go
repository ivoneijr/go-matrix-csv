package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func echo(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

func invert(w http.ResponseWriter, r *http.Request) {
	response := "invert response"
	fmt.Fprint(w, response)
}

func flatern(w http.ResponseWriter, r *http.Request) {
	response := "flatern response"
	fmt.Fprint(w, response)
}

func sum(w http.ResponseWriter, r *http.Request) {
	response := "sum response"
	fmt.Fprint(w, response)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	response := "multiply response"
	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/invert", invert)
	http.HandleFunc("/flatern", flatern)
	http.HandleFunc("/sum", sum)
	http.HandleFunc("/multiply", multiply)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
