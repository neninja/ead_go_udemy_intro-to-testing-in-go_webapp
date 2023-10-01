package main

import (
	"html/template"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
  _ = app.render(w, r, "home.page.tpl", &TemplateData{})
}

type TemplateData struct {
  IP string
  Date map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {
  parsedTemplate, err := template.ParseFiles("./templates/" + t)
  if err != nil {
    http.Error(w, "bad request", http.StatusBadRequest)
  }

  err = parsedTemplate.Execute(w, td)
  
  return nil
}
