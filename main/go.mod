module main

go 1.15

replace (
	GoMicroserviceDemo/handlers => ../handlers
	GoMicroserviceDemo/models => ../models
)

// replace GoMicroserviceDemo/handlers => ../handlers
// replace GoMicroserviceDemo/models => ../models

require (
	GoMicroserviceDemo/handlers v0.0.0-00010101000000-000000000000 // indirect
	GoMicroserviceDemo/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	go.mongodb.org/mongo-driver v1.4.5 // indirect
)
