package router

import (
	"net/http"

	"api/internal/handlers"
	"api/internal/middleware"
	"api/internal/services"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

func New(authService *services.AuthService) http.Handler {
	r := chi.NewRouter()

	r.Use(chimw.Recoverer)
	r.Use(middleware.Logger)

	authHandler := handlers.NewAuthHandler(authService)

	r.Get("/health", handlers.Health)

	r.Route("/api/v1", func(api chi.Router) {
		api.Post("/auth/register", authHandler.Register)
		api.Post("/auth/login", authHandler.Login)

		api.Group(func(protected chi.Router) {
			protected.Use(middleware.JWTAuth(authService))
			protected.Get("/profile", handlers.Profile)
		})
	})

	return r
}
