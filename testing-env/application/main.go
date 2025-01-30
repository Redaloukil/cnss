package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func get_data() ([]byte) {

	resp, err := http.Get("http://localhost:8090/api/v1/data")
	if err != nil {
		log.Println("\n\nserver is down", err)
		return nil 
	} 
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("can't read Response Body:", err)
		return nil
	}
	return body
}


func process_json_data(body []byte ) (map[string]interface{}){
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("can't process json data", err)
		return nil 
	}
	return data
}

func dashboard(w http.ResponseWriter, resp *http.Request){
	fmt.Fprintf(w, "<h1>dashboard<h1>")
}


func main() {
	fmt.Println("\n\nclient running..")

	data := process_json_data(get_data())
	fmt.Println(data)

	http.HandleFunc("/dashboard", dashboard)
	http.ListenAndServe(":8080", nil)
}
