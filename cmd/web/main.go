package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct{
    Session *scs.SessionManager
}

func main() {
	// setup
	app := application{}

	// routes
	mux := app.routes()

    app.Session = getSession()

	// print out a message
	log.Println("Starting server on port 8000...")

	// start server
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
