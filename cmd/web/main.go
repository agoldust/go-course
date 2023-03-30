package main

import (
	"log"
	"net/http"

	"github.com/agoldust/go-course/pkg/config"
	handler "github.com/agoldust/go-course/pkg/handlers"
	"github.com/agoldust/go-course/pkg/render"
)

const portNumber = ":8080"

// main main function of app
func main() {
	// getting Template Cache form disk
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache!")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)

	log.Println("Starting Application on http://localhost:8080")
	_ = http.ListenAndServe(portNumber, nil)
}
