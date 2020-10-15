package main

import (
	"backend/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func MakeRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/signup", handlers.SignUp)
			r.Post("/signup", handlers.SignUp)
		})
	})

	return r
}