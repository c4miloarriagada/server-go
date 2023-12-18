package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"serverpackage/internal/auth"
	db "serverpackage/internal/database"
	"serverpackage/internal/models"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/markbates/goth/gothic"
)

type contextKey string

const AuthProviderKey contextKey = "provider"

var redirectURL string

func GetAuthCallback(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {

		http.Error(w, "Error loading .env", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	provider := chi.URLParam(r, string(AuthProviderKey))

	r = r.WithContext(context.WithValue(context.Background(), AuthProviderKey, provider))

	user, err := gothic.CompleteUserAuth(w, r)

	newUser := &models.User{
		Name:      user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: time.Now(),
		Image:     user.AvatarURL,
	}

	if err := db.DB.FirstOrCreate(&newUser, &models.User{Email: user.Email}).Error; err != nil {
		http.Error(w, "Failed to create or find user", http.StatusInternalServerError)
		log.Fatal(err)
		return

	}

	tokenString, err := auth.CreateToken(user.Email)

	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
	}

	if err != nil {
		http.Error(w, "Failed to create jwt", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	redirectURL := os.Getenv("REDIRECT_URL")

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, redirectURL, http.StatusFound)

}

func BeginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, string(AuthProviderKey))

	r = r.WithContext(context.WithValue(context.Background(), AuthProviderKey, provider))

	gothic.BeginAuthHandler(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
