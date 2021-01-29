package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Application is a handler to handle CRUD requests
type Application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	labs    *LabCollection
}

// func (app *Application) routes() *mux.Router {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/labs", app.getAll).Methods("GET")
// 	return router
// }

// func (app *Application) getAll(rw http.ResponseWriter, r *http.Request) {
// 	list, err := app.labs.GetAll()
// 	if err != nil {
// 		app.serverError(rw, err)
// 	}
// 	err = json.NewEncoder(rw).Encode(list)
// 	if err != nil {
// 		app.serverError(rw, err)
// 	}
// }

// func (app *Application) serverError(rw http.ResponseWriter, err error) {
// 	app.errLog.Println(err)

// 	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// }

// func (app *Application) clientError(w http.ResponseWriter, status int) {
// 	http.Error(w, http.StatusText(status), status)
// }

// // LabCollection represent a mongodb session with a lab data model
// type LabCollection struct {
// 	coll *mongo.Collection
// }

// // GetAll retrieves all labs data
// func (labs *LabCollection) GetAll() ([]Lab, error) {
// 	ctx := context.TODO()
// 	list := []Lab{}

// 	cursor, err := labs.coll.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = cursor.All(ctx, &list)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return list, err
// }

// // A Lab represents lab metadata
// type Lab struct {
// 	ID          primitive.ObjectID `bson:"_id,omitempty"`
// 	Name        string             `bson:"name,omitempty"`
// 	Description string             `bson:"description,omitempty"`
// 	CreatedOn   time.Time          `bson:"createdon,omitempty"`
// }

const (
	serverAddr = "localhost"
	serverPort = 8080
	mongoURL   = "mongodb://127.0.0.1:27017"
	database   = "labs"
	collection = "test"
)

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		errLog.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		errLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	infoLog.Printf("Database connection established")

	// Initialize a new instance of application
	app := &Application{
		infoLog: infoLog,
		errLog:  errLog,
		labs: &LabCollection{
			coll: client.Database(database).Collection(collection),
		},
	}

	// Initialize a new http.Server
	serverURI := fmt.Sprintf("%s:%d", serverAddr, serverPort)

	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}
