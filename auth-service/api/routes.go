package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/verifyToken", app.verifyToken)
	mux.Post("/signup", app.signup)
	mux.Post("/login", app.login)
	mux.Get("/logout", app.logout)

	return mux
}
