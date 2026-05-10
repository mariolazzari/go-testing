package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// app config
	app := application{}

	// app routes
	mux := app.routes()

	// print message
	log.Println("Starting server on port 8080")

	// start server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
