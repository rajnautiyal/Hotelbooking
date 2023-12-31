package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rajnautiyal/bookings/internal/config"
	"github.com/rajnautiyal/bookings/internal/handler"
)

func routes(app *config.AppConfig) http.Handler {
	/*
		this is the router using the Pat
		*mux := pat.New()

		mux.Get("/", http.HandlerFunc(handler.Repo.Home))
		mux.Get("/home", http.HandlerFunc(handler.Repo.Home))
		mux.Get("/about", http.HandlerFunc(handler.Repo.About))*/

	//this is the router using the chi

	mux := chi.NewRouter()

	//loadinf the middleWare for handing common part of request
	mux.Use(middleware.Recoverer)
	mux.Use(NoSruf)
	mux.Use(LoadSession)

	//loading the url for

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/home", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	mux.Get("/generals-quarters", http.HandlerFunc(handler.Repo.General))
	mux.Get("/majors-suite", http.HandlerFunc(handler.Repo.Majors))
	mux.Get("/search-availability", http.HandlerFunc(handler.Repo.Search))
	mux.Post("/search-availability", http.HandlerFunc(handler.Repo.PostSearch))
	mux.Post("/search-availability-json", http.HandlerFunc(handler.Repo.SearchAvaibliltyJson))
	mux.Get("/contact", http.HandlerFunc(handler.Repo.Contact))
	mux.Get("/make-reservation", http.HandlerFunc(handler.Repo.Reservation))
	mux.Post("/make-reservation", http.HandlerFunc(handler.Repo.PostReservation))
	mux.Get("/reservation-summary", http.HandlerFunc(handler.Repo.ReservationSummary))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
