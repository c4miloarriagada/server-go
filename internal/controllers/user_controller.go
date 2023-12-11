package controllers

import (
	"context"
	"fmt"
	"net/http"
	db "serverpackage/internal/database"
	"serverpackage/internal/models"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

type contextKey string

const AuthProviderKey contextKey = "provider"

func GetAuthCallback(w http.ResponseWriter, r *http.Request) {

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
		fmt.Println(w, err)
		return
	}

	if err != nil {
		fmt.Fprintln(w, r)
		return
	}

	http.Redirect(w, r, "http://localhost:5173/home", http.StatusFound)

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
