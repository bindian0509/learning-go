package handlers

import (
	"fmt"
	"net/http"

	"github.com/bindian0509/learning-go/bookings/pkg/config"
	"github.com/bindian0509/learning-go/bookings/pkg/models"
	"github.com/bindian0509/learning-go/bookings/pkg/render"
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
	// adding remote ip to the session 
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, "home.page.htm", r, &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hey Hey Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.htm", r, &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.htm", r, &models.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.htm", r, &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.htm", r, &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.htm", r, &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.htm", r, &models.TemplateData{})
}

func (m *Repository) PostAvailability (w http.ResponseWriter, r *http.Request) {
	
	start_date := r.Form.Get("start")
	end_date := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("So you want the availability from %s to %s", start_date, end_date)))


}
