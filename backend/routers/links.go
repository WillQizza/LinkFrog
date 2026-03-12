package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/willqizza/linkfrog/backend/middleware"
	"github.com/willqizza/linkfrog/backend/models"
	"github.com/willqizza/linkfrog/backend/services"
	"github.com/willqizza/linkfrog/backend/utils"
)

func linksRouter() chi.Router {
	router := chi.NewRouter()

	router.With(middleware.AuthRequired).Get("/", handleGetLinks)
	router.With(middleware.AuthRequired).Post("/link", handleCreateLink)
	router.With(middleware.AuthRequired).Delete("/link/{path}", handleDeleteLink)
	router.Get("/redirect/{path}", handleRedirect)

	return router
}

func handleGetLinks(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)

	links, err := services.GetLinksByUser(r.Context(), user)
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "An error occurred while attempting to retrieve your links"})
		return
	}

	utils.WriteJSON(w, 200, map[string]any{
		"links": links,
	})
}

func handleCreateLink(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)

	var payload struct {
		Link string
		Code *string
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteJSON(w, 400, map[string]string{"error": "Invalid Payload"})
		return
	}

	var code string
	goodCode := false
	if payload.Code != nil {
		codeExists, err := services.DoesCodeExist(r.Context(), *payload.Code)
		if err != nil {
			utils.WriteJSON(w, 500, map[string]string{"error": "An error occurred while attempting to retrieve existing codes"})
			return
		}

		if codeExists {
			utils.WriteJSON(w, 409, map[string]string{"error": "The code requested already exists"})
			return
		}

		code = *payload.Code
		goodCode = true
	} else {
		for range 10 {
			code = utils.RandomCode()

			codeExists, err := services.DoesCodeExist(r.Context(), code)
			if err != nil {
				utils.WriteJSON(w, 500, map[string]string{"error": "An error occurred while attempting to retrieve existing codes"})
				return
			}

			if !codeExists {
				goodCode = true
				break
			}
		}
	}

	if !goodCode {
		utils.WriteJSON(w, 500, map[string]string{"error": "Unable to autogenerate a code. Please try again later."})
		return
	}

	err = services.CreateLink(r.Context(), &models.Link{
		Owner: user.ID,
		Path:  code,
		URL:   payload.Link,
	})
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "Unable to create link. Please try again later."})
		return
	}

	utils.WriteJSON(w, 200, map[string]string{
		"message": "Successfully shortened link",
		"code":    code,
	})
}

func handleDeleteLink(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	path := chi.URLParam(r, "path")

	link, err := services.GetLinkByCode(r.Context(), path)
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "An error occurred while attempting to retrieve the link"})
		return
	}

	if link == nil {
		utils.WriteJSON(w, 404, map[string]string{"error": "Link not found"})
		return
	}

	if link.Owner != user.ID {
		utils.WriteJSON(w, 403, map[string]string{"error": "You do not have permission to delete this link"})
		return
	}

	err = services.DeleteLink(r.Context(), path)
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "An error occurred while attempting to delete the link"})
		return
	}

	utils.WriteJSON(w, 200, map[string]string{"message": "Successfully deleted link"})
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "path")

	link, err := services.GetLinkByCode(r.Context(), path)
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "An error occurred while attempting to retrieve the link"})
		return
	}

	if link == nil {
		utils.WriteJSON(w, 404, map[string]string{"error": "Link not found"})
		return
	}

	http.Redirect(w, r, link.URL, http.StatusFound)
}
