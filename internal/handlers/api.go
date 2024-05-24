package handlers

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jantoniogonzalez/go-web-server/internal/middleware"
)

func Handler(r chi.Router) {
	// Global Middleware
	r.Use(chiMiddleware.StripSlashes)

	r.Route("/account", func(r chi.Router) {

		r.Use(middleware.Authorization)
		r.Get("/coins", GetCoinBalance) // GET user account balance
	})
}
