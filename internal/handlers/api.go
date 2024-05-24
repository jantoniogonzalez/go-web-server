package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler(r chi.Router) {
	// Global Middleware
	r.Use(middleware.StripSlashes)
}
