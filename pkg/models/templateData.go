package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRToken  string
	Flash     string
	WARN      string
	Error     string
}
