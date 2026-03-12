package middleware

import (
	"context"
	"net/http"

	"github.com/willqizza/linkfrog/backend/services"
	"github.com/willqizza/linkfrog/backend/utils"
)

const UserKey = "user"

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtCookie, err := r.Cookie("token")
		if err != nil {
			utils.WriteJSON(w, 401, map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		userId, err := utils.ParseJWT(jwtCookie.Value)
		if err != nil {
			utils.WriteJSON(w, 401, map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		user, err := services.GetUserByID(r.Context(), userId)
		if err != nil {
			utils.WriteJSON(w, 401, map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
