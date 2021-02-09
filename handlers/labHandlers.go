package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"GoMicroserviceDemo/models"

	"github.com/gorilla/mux"
)

// Application is a handler to handle CRUD requests
type Application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	labs    *models.LabCollection
}

// NewApplication creates a new handler
func NewApplication(info *log.Logger, err *log.Logger, labs *models.LabCollection) *Application {
	return &Application{info, err, labs}
}

// All calls GetAll func from labs to retrive all labs info from database
func (app *Application) All(rw http.ResponseWriter, r *http.Request) {
	list, err := app.labs.GetAll()
	if err != nil {
		app.serverError(rw, err)
	}
	err = json.NewEncoder(rw).Encode(list)
	if err != nil {
		app.serverError(rw, err)
	}
}

// Create calls CreateLab func from labs to insert a new lab into database
func (app *Application) Create(rw http.ResponseWriter, r *http.Request) {
	var lab models.Lab
	err := json.NewDecoder(r.Body).Decode(&lab)
	if err != nil {
		app.serverError(rw, err)
	}
	lab.CreatedOn = time.Now()
	insert, err := app.labs.CreateLab(lab)
	if err != nil {
		app.serverError(rw, err)
	}
	// app.infoLog.Printf("Created a new lab, id=%s", insert.InsertedID)
	app.infoLog.Printf("Created a new lab: %s", insert)
}

// UpdateName calls UpdateLabName func from labs to update a lab name in database
func (app *Application) UpdateName(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	newName := vars["name"]
	updatedDoc, err := app.labs.UpdateLabName(id, newName)
	if err != nil {
		app.serverError(rw, err)
	}
	app.infoLog.Printf("%s has been updated", updatedDoc)
}

// Delete calls DeleteLab func from labs to delete a lab in database
func (app *Application) Delete(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	deleteRes, err := app.labs.DeleteLab(id)
	if err != nil {
		app.serverError(rw, err)
	}
	app.infoLog.Printf("Deleted %d lab", deleteRes.DeletedCount)
}
