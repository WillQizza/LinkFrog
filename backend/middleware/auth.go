package middleware

import (
	"context"
	"net/http"

	"github.com/willqizza/linkfrog/backend/utils"
)

const UserIDKey = "userId"

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtCookie, err := r.Cookie("token")
		if err != nil {
			utils.WriteJSON(w, 401, map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		token := jwtCookie.Value

		userId, err := utils.ParseJWT(token)
		if err != nil {
			utils.WriteJSON(w, 401, map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
