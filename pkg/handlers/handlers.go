package handler

import (
	"log"
	"net/http"

	"github.com/agoldust/go-course/pkg/render"
)

// Home this is about Home page handeler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
	log.Println("Home Page Rendered!")
}

// Home this is about about page handeler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
	log.Println("About Page Rendered!")
}
