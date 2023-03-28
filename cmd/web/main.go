package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

// main main function of app
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting Application on http://localhost:8080")
	_ = http.ListenAndServe(portNumber, nil)
}
