package routers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/willqizza/linkfrog/backend/middleware"
)

func linksRouter() chi.Router {
	router := chi.NewRouter()

	router.With(middleware.AuthRequired).Get("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Grab all of the links we've created
	})

	router.With(middleware.AuthRequired).Post("/shorten", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Create a new link shortened
		var userId int = r.Context().Value(middleware.UserIDKey).(int)
		fmt.Println(userId)

		w.Write([]byte("test"))
	})

	router.Get("/redirect/{path}", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Redirects based off of the path
	})

	return router
}
