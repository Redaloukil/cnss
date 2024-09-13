package main

// TODO: recive interface return struct
import (
	"log"
	"net/http"
	"html/template"
)


func index(w http.ResponseWriter, r *http.Request)  {
	tmpl := template.Must(template.ParseFiles("./components/index.html")) 
	tmpl.Execute(w, nil)
}


func main() {
	log.Println("Server is running on http://localhost:8080")

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
