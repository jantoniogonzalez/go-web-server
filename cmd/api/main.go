package main

import (
	"fmt"
	"net/http"

	"github.com/jantoniogonzalez/go-web-server/internal/handlers"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request)) {
		w.write([byte("welcome")])
	}
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Error(err)
	}
}