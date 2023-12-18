package server

import (
	"context"
	"net/http"
	"serverpackage/internal/controllers"
	"serverpackage/internal/middlewares"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use((middleware.Logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//auth
	r.Get("/auth/{provider}", s.beginAuthProviderCallback)
	r.Get("/auth/{provider}/callback", controllers.GetAuthCallback)
	r.Get("/logout/{provider}", controllers.Logout)

	//
	r.With(middlewares.JwtHandler).Get("/getUser", controllers.GetUser)

	return r

}

func (s *Server) beginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}
