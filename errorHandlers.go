package main

import "net/http"

func (app *Application) serverError(rw http.ResponseWriter, err error) {
	app.errLog.Println(err)

	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
