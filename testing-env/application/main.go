package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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

func data(w http.ResponseWriter, r *http.Request)  {

	url := "http://localhost:8090/api/v1/data"

	// get url json response
	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w,"error while feching...\n ")
		fmt.Fprintf(w,"%v",err)
		return
	}
	defer response.Body.Close()

	// check status code 
	if response.StatusCode != http.StatusOK {
		fmt.Fprintf(w,"\nNon-OK HTTP status")
		fmt.Fprintf(w,"%v\n",response.StatusCode)
		return
	}

	// read slice byte data 
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w,"Error reading response body: %v", err)
		return
	}


	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Fprintf(w,"Error while processing json : %v", err)
		return
	}

	fmt.Fprintf(w, "%v\n", result)
	fmt.Fprintf(w, "\nGot respnce")
}


func main() {
	log.Println("Server is running on http://localhost:8080")

	http.HandleFunc("/", index)
	http.HandleFunc("/data", data)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

