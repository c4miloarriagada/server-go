package server

import (
	"context"
	"net/http"
	"serverpackage/internal/controllers"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use((middleware.Logger))
	r.Get("/auth/{provider}", middleware.JwtHandler(s.beginAuthProviderCallback))
	r.Get("/auth/{provider}/callback", controllers.GetAuthCallback)
	r.Get("/logout/{provider}", controllers.Logout)

	return r

}

func (s *Server) beginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}
