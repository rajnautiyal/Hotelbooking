package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rajnautiyal/bookings/pkg/config"
	"github.com/rajnautiyal/bookings/pkg/handler"
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

	mux.Use(middleware.Recoverer)
	mux.Use(NoSruve)
	mux.Use(LoadSession)
	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/home", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	return mux
}
