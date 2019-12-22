package main

import (
	"log"
	"net/http"

	"./routes"
)

// main function to handle routes
func main() {
	http.HandleFunc("/echo", routes.Echo)
	http.HandleFunc("/invert", routes.Invert)
	http.HandleFunc("/flatten", routes.Flatten)
	http.HandleFunc("/sum", routes.Sum)
	http.HandleFunc("/multiply", routes.Multiply)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
