package titles

import (
	"github.com/go-chi/chi/v5"
)

func LoadTitleRoutes(mainRouter *chi.Mux) {

	r := chi.NewRouter()

	r.Get("/{id:[1-9]\\d*}", getTitle)
	r.Get("/", getTitles)

	mainRouter.Mount("/titles", r)
}
