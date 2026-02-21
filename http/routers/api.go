package routers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func apiRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome api"))
	})

	return router
}
