package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func get_data(){

	resp, err := http.Get("http://localhost:8090/api/v1/data")
	if err != nil {
		log.Println("\n\nserver is down", err)
		return
	} 
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println("\n\ncan't read Response body", err)
		return
    }
	
}


func main() {
	fmt.Println("\n\nclient running..")
	get_data()
}
