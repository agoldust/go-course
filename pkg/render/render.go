package render

import (
	"log"
	"net/http"
	"text/template"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println("Error parsing the template")
		return
	}

}