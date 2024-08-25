package main

// TODO: recive interface return struct
import (
	"encoding/json"
	"fmt"
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

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/data", data)
	http.HandleFunc("/user", data)
	http.ListenAndServe(":8080", nil)
}
