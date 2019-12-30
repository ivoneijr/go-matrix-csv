package main

import (
	"log"
	"net/http"

	"./handlers"
)

// main function to handle routes
func main() {
	http.HandleFunc("/health-check", handlers.HealthCheckHandler)
	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/invert", handlers.Invert)
	http.HandleFunc("/flatten", handlers.Flatten)
	http.HandleFunc("/sum", handlers.Sum)
	http.HandleFunc("/multiply", handlers.Multiply)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
