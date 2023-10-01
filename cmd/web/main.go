package main

import (
	"log"
	"net/http"
)

type application struct {}

func main() {
  // setup
  app := application{}

  // routes
  mux := app.routes()

  // print out a message
  log.Println("Starting server on port 8000...")

  // start server
  err := http.ListenAndServe(":8000", mux)
  if err != nil {
    log.Fatal(err)
  }
}
