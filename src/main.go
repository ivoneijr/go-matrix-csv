package main

import (
	"log"
	"net/http"

	"./routes"
)

func main() {
	http.HandleFunc("/echo", routes.Echo)
	http.HandleFunc("/flatten", routes.Flatten)
	http.HandleFunc("/sum", routes.Sum)
	http.HandleFunc("/multiply", routes.Multiply)
	http.HandleFunc("/invert", routes.Invert)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
