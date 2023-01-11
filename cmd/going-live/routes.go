package main

import (
	"github.com/go-chi/chi/v5"

	"github.com/advance467/Going-Live/cmd/going-live/features/titles"
)

func MainRoutes() *chi.Mux {
	r := chi.NewRouter()
	titles.LoadTitleRoutes(r)
	return r
}
