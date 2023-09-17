package routes

import (
	"auth-service/internal/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	http.Handler
	service user.Service
}

func NewRouter(service user.Service) Router {
	router := Router{
		service: service,
	}
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/signup", router.signup)
	mux.Post("/login", router.login)
	mux.Put("/{userID}", router.updateUser)
	mux.Delete("/{userID}", router.deleteUser)
	mux.Put("/{userID}/password", router.updatePassword)

	router.Handler = mux
	return router
}
