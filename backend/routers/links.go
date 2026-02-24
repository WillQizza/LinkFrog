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

	return router
}
