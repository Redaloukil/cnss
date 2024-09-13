package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed components/*
var content embed.FS

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmpls, err := template.ParseFS(content, "components/*.html")
	if err != nil {
		http.Error(w, "could not parse template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	err = tmpls.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		http.Error(w, "could not execute template", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func main() {
	log.Println("Server is running on http://localhost:8080")

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

