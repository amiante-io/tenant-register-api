package main

import (
	"encoding/json"
	"net/http"
)

type Tenant struct {
	Name string `json:"name"`
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	js, err := json.MarshalIndent(Tenant{Name: "tenant-test"}, "", "\n")
	if err != nil {
		// todo handle error
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
