package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var url string = "http://localhost:8090"

func get_data() map[string]interface{} {

	resp, err := http.Get(url + "/api/v1/data")
	if err != nil {
		log.Println("\n\nserver is down", err)
		return nil
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("can't read Response Body:", err)
		return nil
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("can't process json data", err)
		return nil
	}
	fmt.Println("Response status:", resp.Status)
	return data
}

func dashboard()  {
	data := get_data()

	fmt.Println("\n*---------entry-----------")
	for key, value := range data{
		fmt.Println("|",key, ":", value)
	}
	fmt.Println("*------------------------------")
}

func main(){
	dashboard()
}
