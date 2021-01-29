package main

import "github.com/gorilla/mux"

func (app *Application) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/labs", app.getAll).Methods("GET")
	return router
}
