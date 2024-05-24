package api

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jantoniogonzalez/go-web-server/internal/middleware"
)

func Handler(r chi.Router) {
	// Global Middleware
	r.Use(chiMiddleware.StripSlashes)

	// // A normal middleware stack, but i think they are mainly for logging purposes
	// r.Use(chiMiddleware.RequestID) // Request Id to header
	// r.Use(chiMiddSleware.RealIP) // Only use if you trust the headers being passed
	// r.Use(chiMiddleware.Logger) // I think does the same as logrus??
	// r.Use(chiMiddleware.Recoverer) // errorrr

	r.Route("/account", func(r chi.Router) {

		r.Use(middleware.Authorization)
		r.Get("/coins", GetCoinBalance) // GET user account balance
	})
}
