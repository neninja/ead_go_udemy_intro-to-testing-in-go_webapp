package main

import (
	"flag"
	"log"
	"net/http"
	"webapp/pkg/db"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	DSN     string
	DB      db.PostgresConn
	Session *scs.SessionManager
}

func main() {
	// setup
	app := application{}

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=ead_go_udemy_intro-to-testing-in-go_webapp sslmode=disable timezone=UTC connect_timeout=5", "Postgress connection")
	flag.Parse()

	conn, err := app.connectToDB()
    if err != nil {
        log.Fatal(err)
    }

    app.DB = db.PostgresConn{DB: conn}

	app.Session = getSession()

	// print out a message
	log.Println("Starting server on port 8000...")

	// start server
	err = http.ListenAndServe(":8000", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
