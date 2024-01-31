package main

import (
	"net/http"
)

type Tenant struct {
	Name string `json:"name"`
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	err := app.writeJson(w, http.StatusOK, &Tenant{Name: "tenant-test"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
