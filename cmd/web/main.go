package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rajnautiyal/bookings/internal/config"
	"github.com/rajnautiyal/bookings/internal/handler"
	"github.com/rajnautiyal/bookings/internal/helpers"
	"github.com/rajnautiyal/bookings/internal/models"
	"github.com/rajnautiyal/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	//

	/*
		http.HandleFunc("/", handler.Repo.Home)
		http.HandleFunc("/about", handler.Repo.About)
	*/
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	fmt.Printf("Staring application on port %s", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
func run() error {
	gob.Register(models.Reservation{})
	app.InProduction = false
	infoLog = log.New(os.Stdout, "INFO/t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "Error/t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Issue while creating the template")
		return err
	}
	render.NewTemplate(&app)

	app.UseCache = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	repo := handler.NewRep(&app)
	handler.NewHandler(repo)
	helpers.NewHelpers(&app)
	if err != nil {
		log.Fatal("Issue while creating teh handler")
		return err
	}
	app.TemplateCache = tc
	return nil
}
