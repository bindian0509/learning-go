package handlers

import (
	"net/http"

	"github.com/bindian0509/learning-go/web-app-html/pkg/config"
	"github.com/bindian0509/learning-go/web-app-html/pkg/models"
	"github.com/bindian0509/learning-go/web-app-html/pkg/render"
)

// Repository used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository for handlers
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the new repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.htm", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hey Hey Hello, again"

	// send data to the template
	render.RenderTemplate(w, "about.page.htm", &models.TemplateData{
		StringMap: stringMap,
	})
}
