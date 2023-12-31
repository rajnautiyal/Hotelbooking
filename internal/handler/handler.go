package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rajnautiyal/bookings/internal/config"
	"github.com/rajnautiyal/bookings/internal/forms"
	"github.com/rajnautiyal/bookings/internal/helpers"
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
	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{}, request)
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{}, request)
}

func (m *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, request)
}

// PostReservation handles the posting of a reservation form
// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	reservation := models.Reservation{
		FirstName: request.Form.Get("first_name"),
		LastName:  request.Form.Get("last_name"),
		Email:     request.Form.Get("email"),
		Phone:     request.Form.Get("phone"),
	}

	form := forms.New(request.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, request)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		}, request)
		return
	}
	fmt.Println("redirecting to summary page ")
	m.App.Session.Put(request.Context(), "reservation", reservation)
	http.Redirect(writer, request, "/reservation-summary", http.StatusSeeOther)

}

func (m *Repository) ReservationSummary(writer http.ResponseWriter, request *http.Request) {
	reservation, ok := m.App.Session.Get(request.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("can't get the server error")
		m.App.Session.Put(request.Context(), "error", "no reservation session value")
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
	}
	//m.App.Session.Remove(request.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(writer, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	}, request)
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

	out, err := json.MarshalIndent(availabilityJson, "", "")
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(out)
}
