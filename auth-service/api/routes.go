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

	mux.NotFound(app.invalidPath)

	mux.Get("/verifyToken", app.verifyToken)
	mux.Post("/signup", app.signup)
	mux.Post("/login", app.login)

	mux.Put("/{userID}", app.updateUser)
	mux.Put("/{userID}/password", app.updatePassword)
	mux.Get("/{userID}/logout", app.logout)
	mux.Delete("/{userID}", app.deleteUser)

	return mux
}
