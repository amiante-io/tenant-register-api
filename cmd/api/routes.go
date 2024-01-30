package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", app.register)

	return r
}
