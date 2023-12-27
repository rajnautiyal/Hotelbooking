package render

import (
	"bytes"
	"fmt"

	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/rajnautiyal/bookings/internal/config"
	"github.com/rajnautiyal/bookings/internal/models"
)

var app *config.AppConfig

// set the config for the
func NewTemplate(config *config.AppConfig) {
	app = config
}

func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
	templateData.CSRFToken = nosurf.Token(request)
	fmt.Println("adding SSFRtomeoke", templateData.CSRFToken)
	return templateData
}
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData, request *http.Request) {

	var tc map[string]*template.Template
	//create the template  cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get the template")
	}

	buf := new(bytes.Buffer)
	templateData = AddDefaultData(templateData, request)
	_ = t.Execute(buf, templateData)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	//get all the file ends with page.tmpl
	pages, err := filepath.Glob("./templates/*page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

/*

// old way to rendering the table after creating the map

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	_, inMap := tc[t]
	log.Println("this is my home template", t)
	if !inMap {
		log.Println("creating a template and storing in the logs")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		log.Println("using the cach function")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("this is some error", err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
*/
