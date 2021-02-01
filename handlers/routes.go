package handlers

import "github.com/gorilla/mux"

// Routes returns a router matching incoming requests to their respective handler
func (app *Application) Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/labs", app.All).Methods("GET")
	router.HandleFunc("/newLab", app.Create).Methods("POST")
	router.HandleFunc("/labs/{id}/{name}", app.UpdateName).Methods("PUT")
	router.HandleFunc("/labs/{id}", app.Delete).Methods("DELETE")
	return router
}
