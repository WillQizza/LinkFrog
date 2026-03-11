package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/willqizza/linkfrog/backend/middleware"
	"github.com/willqizza/linkfrog/backend/utils"
)

func linksRouter() chi.Router {
	router := chi.NewRouter()

	router.With(middleware.AuthRequired).Get("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Grab all of the links we've created
	})

	router.With(middleware.AuthRequired).Post("/link", func(w http.ResponseWriter, r *http.Request) {
		// userId := r.Context().Value(middleware.UserIDKey).(int)

		var payload struct {
			link string
			code *string
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.WriteJSON(w, 400, map[string]string{
				"error": "Invalid Payload",
			})
			return
		}

		// TODO: Verify link is valid
		utils.WriteJSON(w, 200, map[string]string{
			"message": "Successfully shortened link",
		})
	})

	router.With(middleware.AuthRequired).Delete("/link/{path}", func(w http.ResponseWriter, r *http.Request) {

	})

	router.Get("/redirect/{path}", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Redirects based off of the path
	})

	return router
}
