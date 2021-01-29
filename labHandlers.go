package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) getAll(rw http.ResponseWriter, r *http.Request) {
	list, err := app.labs.GetAll()
	if err != nil {
		app.serverError(rw, err)
	}
	err = json.NewEncoder(rw).Encode(list)
	if err != nil {
		app.serverError(rw, err)
	}
}
