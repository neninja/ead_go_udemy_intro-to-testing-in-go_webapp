package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"
)

var pathToTemplates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
    td := make(map[string]any)

    if app.Session.Exists(r.Context(), "test") {
        msg := app.Session.GetString(r.Context(), "test")
        td["test"] = msg
    } else {
        app.Session.Put(r.Context(), "test", "Hit this page at " + time.November.String())
    }

    _ = app.render(w, r, "home.page.tpl", &TemplateData{Data: td})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse the template from disk.
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "base.layout.tpl"))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	// execute the template, passing it data, if any
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	form := NewForm(r.PostForm)
	form.Required("email", "password")

	if !form.Valid() {
		fmt.Fprintf(w, "failed validation")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	fmt.Println(email, password)

	fmt.Fprint(w, email)
}
