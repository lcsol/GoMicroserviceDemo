package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"GoMicroserviceDemo/handlers"
	"GoMicroserviceDemo/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// // Application is a handler to handle CRUD requests
// type Application struct {
// 	infoLog *log.Logger
// 	errLog  *log.Logger
// 	labs    *LabCollection
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
	collection := models.NewLabCollection(client.Database(database).Collection(collection))
	app := handlers.NewApplication(infoLog, errLog, collection)

	// Initialize a new http.Server
	serverURI := fmt.Sprintf("%s:%d", serverAddr, serverPort)

	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}
