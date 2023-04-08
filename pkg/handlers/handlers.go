package handlers

import (
	"net/http"

	"github.com/agoldust/go-course/pkg/config"
	"github.com/agoldust/go-course/pkg/modules"
	"github.com/agoldust/go-course/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home this is about Home page handeler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["page-title"] = "Home"

	// adding test data to sessions
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// rendering the template
	render.RenderTemplate(w, "home.page.tmpl", &modules.TemplateData{
		StringMap: stringMap,
	})
}

// Home this is about about page handeler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["page-title"] = "About"

	// adding test data from sessions
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &modules.TemplateData{
		StringMap: stringMap,
	})
}
