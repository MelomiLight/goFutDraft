package main

import (
	_ "github.melomii/futDraft/internal/data"
	data2 "github.melomii/futDraft/internal/data"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (app *application) futDraftHandler(writer http.ResponseWriter, request *http.Request) {
	//tmpl, _ := template.ParseFiles("templates/futdraft.html")
	//tmpl.Execute(writer, data)
	//http.ListenAndServe(":4000", nil)

	fs := http.FileServer(http.Dir(".cmd/api/templates"))
	http.Handle(".cmd/api/templates/", http.StripPrefix(".cmd/api/templates/", fs))
	http.HandleFunc("/", serveTemplate)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	data := data2.Players{
		ID:         123,
		CommonName: "Ronaldu",
	}
	lp := filepath.Join("templates", "futdraft.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Print(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
