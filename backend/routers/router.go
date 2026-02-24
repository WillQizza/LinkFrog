package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Mount("/api/links", linksRouter())
	router.Mount("/api/auth", authRouter())

	return router
}
