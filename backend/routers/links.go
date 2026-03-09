package routers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func linksRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("links api"))
	})

	router.Post("/shorten", func(w http.ResponseWriter, r *http.Request) {

	})

	router.Get("/redirect", func(w http.ResponseWriter, r *http.Request) {

	})

	return router
}
