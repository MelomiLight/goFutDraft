package main

import (
	_ "github.melomii/futDraft/internal/data"
	data2 "github.melomii/futDraft/internal/data"
	"html/template"
	"net/http"
)

func (app *application) futDraftHandler(writer http.ResponseWriter, request *http.Request) {
	player := data2.Players{
		ID:         123,
		CommonName: "Ronaldu",
	}
	//tmpl, _ := template.ParseFiles("templates/futdraft.html")
	//tmpl.Execute(writer, data)
	tmpl := template.Must(template.ParseFiles("./cmd/api/templates/futdraft.html"))

	// Render the template
	err := tmpl.Execute(writer, player)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) futDraftChooseHandler(writer http.ResponseWriter, request *http.Request) {

}
