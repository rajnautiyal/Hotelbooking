package models

import "github.com/rajnautiyal/bookings/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	WARN      string
	Error     string
	Form      *forms.Form

	
}
