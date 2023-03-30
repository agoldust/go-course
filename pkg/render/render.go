package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/agoldust/go-course/pkg/config"
	"github.com/agoldust/go-course/pkg/modules"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *modules.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache { // getting the template cache from App Config
		tc = app.TemplateCache
		log.Println("using cache")
	} else {
		tc, _ = CreateTemplateCache()
		log.Println("creating cache")
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(("Could not get template form template cache"))
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)
	if err != nil {
		log.Panicln(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Panicln(err)
	}
}

// CreateTemplateCache creating the template cache of templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// getting all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files in pages
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// locking for layouts that
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
