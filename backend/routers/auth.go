package routers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/mail"
	"os"

	"github.com/go-chi/chi"
	"github.com/willqizza/linkfrog/backend/middleware"
	"github.com/willqizza/linkfrog/backend/services"
	"github.com/willqizza/linkfrog/backend/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"openid", "email", "profile"},
	Endpoint:     google.Endpoint,
}

type googleApiUserInfo struct {
	Sub   string
	Email string
	Name  string
}

func authRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/google", handleGoogleLogin)
	router.Get("/google/callback", handleGoogleCallback)
	router.With(middleware.AuthRequired).Post("/invite", handleInvite)

	return router
}

func generateState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	state, err := generateState()
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "Failed to generate state"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   300,
	})
	url := googleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Check state cookie and remove it since it's no longer valid.
	stateCookie, err := r.Cookie("oauth_state")

	http.SetCookie(w, &http.Cookie{
		Name:   "oauth_state",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	if err != nil || stateCookie.Value != r.URL.Query().Get("state") {
		utils.WriteJSON(w, 400, map[string]string{"error": "Invalid state"})
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		utils.WriteJSON(w, 400, map[string]string{"error": "Missing code"})
		return
	}

	token, err := googleOAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "Failed to exchange token"})
		return
	}

	client := googleOAuthConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	var googleUser googleApiUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "Failed to decode user info"})
		return
	}

	// Check if user exists or if this is the admin user that needs to be added
	userId, err := services.GetUserIdByEmail(r.Context(), googleUser.Email)
	if err == sql.ErrNoRows {
		// Check if any users exist at all
		count, err := services.GetTotalUsers(r.Context())
		if err != nil {
			utils.WriteJSON(w, 500, map[string]string{"error": "Database error"})
			return
		}

		if count > 0 {
			// User is not whitelisted
			http.Redirect(w, r, os.Getenv("AUTH_REDIRECT_URL")+"?error=unauthorized", http.StatusTemporaryRedirect)
			return
		}

		// First user is an admin user and can be inserted
		userId, err = services.WhitelistUser(r.Context(), googleUser.Email)
		if err != nil {
			utils.WriteJSON(w, 500, map[string]string{"error": "Failed to create user"})
			return
		}
	} else if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "Database error"})
		return
	}

	jwtToken, err := utils.SignJWT(userId)
	if err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "Failed to sign token"})
		return
	}

	http.Redirect(w, r, os.Getenv("AUTH_REDIRECT_URL")+"?token="+jwtToken, http.StatusTemporaryRedirect)
}

func handleInvite(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email string
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil || payload.Email == "" {
		utils.WriteJSON(w, 400, map[string]string{"error": "Invalid payload"})
		return
	}

	if _, err = mail.ParseAddress(payload.Email); err != nil {
		utils.WriteJSON(w, 400, map[string]string{"error": "Invalid payload"})
		return
	}

	if _, err = services.GetUserIdByEmail(r.Context(), payload.Email); err == nil {
		utils.WriteJSON(w, 409, map[string]string{"error": "User already whitelisted"})
		return
	} else if err != sql.ErrNoRows {
		utils.WriteJSON(w, 500, map[string]string{"error": "An error occurred while attempting to check if the email is already whitelisted"})
		return
	}

	if _, err = services.WhitelistUser(r.Context(), payload.Email); err != nil {
		utils.WriteJSON(w, 500, map[string]string{"error": "An error occurred while attempting to whitelist the email"})
		return
	}

	utils.WriteJSON(w, 200, map[string]string{"message": "User invited successfully"})
}
