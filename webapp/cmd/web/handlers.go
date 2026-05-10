package main

import (
	"html/template"
	"net/http"
)

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {

	// parse tempate
	parsed, err := template.ParseFiles("./templates/" + t)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	// exec template
	err = parsed.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", &TemplateData{})
}
