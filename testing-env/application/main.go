package main

// TODO: recive interface return struct
import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)



func getData() (*map[string]interface{}, error) {
	addressURL := "http://localhost:8090/api/v1/data"
	result, err := http.Get(addressURL)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", result.StatusCode)
	}

	var data map[string]interface{}
	err = json.NewDecoder(result.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return &data, nil
}

func index(responseProssess http.ResponseWriter, responce *http.Request) {
	fmt.Fprint(responseProssess, "<h1>this is a sample web application</h1>")
}


func signin(responseProssess http.ResponseWriter, responce *http.Request) {
	fmt.Fprint(responseProssess, "new? click here to sign in <button> sign in</button>")
}

func user(responseProssess http.ResponseWriter, responce *http.Request) {
	fmt.Fprint(responseProssess, "user dash-board")
}


func data(responseProssess http.ResponseWriter, responce *http.Request) {
	data, err := getData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Fprint(responseProssess, data)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "Admin"
	}
	w.Write([]byte("login as, " + name + "!"))
}


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/data", data)
	http.HandleFunc("/user", data)
	http.HandleFunc("/dashboard", handleIndex)
	http.HandleFunc("/login", handleGreet)
	http.HandleFunc("/signin", signin)
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
