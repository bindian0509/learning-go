package handlers

import (
	"net/http"
	"github.com/bindian0509/learning-go/web-app-html/pkg/config"
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
	return &Repository {
		App: a,

	}
}
// NewHandlers sets the new repository for handlers 
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.htm")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.htm")
}
