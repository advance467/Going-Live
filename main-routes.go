package main

import (
	"github.com/advance467/Going-Live/features/titles"
	"github.com/go-chi/chi/v5"
)

func MainRoutes() *chi.Mux {

	r := chi.NewRouter()

	titles.LoadTitleRoutes(r)

	return r

}
