package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rajnautiyal/bookings/internal/config"
	"github.com/rajnautiyal/bookings/internal/handler"
	"github.com/rajnautiyal/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	tc, err := render.CreateTemplateCache()
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
	if err != nil {
		log.Println("error while creating the template")
	}
	app.TemplateCache = tc
	/*
		http.HandleFunc("/", handler.Repo.Home)
		http.HandleFunc("/about", handler.Repo.About)
	*/
	fmt.Printf(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	//fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
