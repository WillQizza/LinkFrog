package routers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func rootRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return router
}
