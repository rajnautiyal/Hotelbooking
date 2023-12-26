package handler

import (
	"net/http"

	"github.com/rajnautiyal/bookings/pkg/config"
	"github.com/rajnautiyal/bookings/pkg/models"
	"github.com/rajnautiyal/bookings/pkg/render"
)

// template data hold data send data from handle to template

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRep(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {

	//perform the logic
	remoteIp := m.App.Session.GetString(request.Context(), "remote_ip")
	myMap := make(map[string]string)
	myMap["test"] = "hello Again..."
	myMap["remoteIp"] = remoteIp
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: myMap,
	})
}
