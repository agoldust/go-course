package main

import (
	"log"
	"net/http"

	handler "github.com/agoldust/go-course/pkg/handlers"
)

const portNumber = ":8080"

// main main function of app
func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	log.Println("Starting Application on http://localhost:8080")
	_ = http.ListenAndServe(portNumber, nil)
}
