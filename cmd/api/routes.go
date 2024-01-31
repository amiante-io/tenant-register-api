package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.NotFound(app.notFoundResponse)

	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Get("/", app.register)

	return r
}
