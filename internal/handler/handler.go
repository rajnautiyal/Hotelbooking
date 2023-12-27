package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rajnautiyal/bookings/internal/config"
	"github.com/rajnautiyal/bookings/internal/forms"
	"github.com/rajnautiyal/bookings/internal/models"
	"github.com/rajnautiyal/bookings/internal/render"
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
	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{}, request)
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {

	//perform the logic
	remoteIp := m.App.Session.GetString(request.Context(), "remote_ip")
	myMap := make(map[string]string)
	myMap["test"] = "hello Again..."
	myMap["remoteIp"] = remoteIp
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: myMap,
	}, request)
}

func (m *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	}, request)
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: request.Form.Get("first_name"),
		LastName:  request.Form.Get("last_name"),
		Email:     request.Form.Get("email"),
		Phone:     request.Form.Get("phone"),
	}

	form := forms.New(request.PostForm)

	form.Has("first_name", request)

	if !form.Valid() {

		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		}, request)
		return
	}
}

func (m *Repository) Majors(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "majors.page.tmpl", &models.TemplateData{}, request)
}

func (m *Repository) General(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "generals.page.tmpl", &models.TemplateData{}, request)
}

func (m *Repository) Search(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "search-availability.page.tmpl", &models.TemplateData{}, request)
}

func (m *Repository) PostSearch(writer http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")
	writer.Write([]byte(fmt.Sprintf("retrive the search data from %s to %s", start, end)))
}
func (m *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "contact.page.tmpl", &models.TemplateData{}, request)
}

type JsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) SearchAvaibliltyJson(writer http.ResponseWriter, request *http.Request) {
	availabilityJson := JsonResponse{
		Ok:      true,
		Message: "Available",
	}

	fmt.Println("I am not here ")
	out, err := json.MarshalIndent(availabilityJson, "", "")
	if err != nil {
		log.Fatalf(err.Error())
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(out)
}
