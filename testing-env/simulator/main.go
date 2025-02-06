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

func dashboard(w http.ResponseWriter, resp *http.Request)  {
	data := get_data()

	fmt.Fprintf(w, "%v", "\n*---------entry-----------")
	for key, value := range data{
		fmt.Fprintf(w, "\n| %v : %v", key, value)
	}
	fmt.Fprintf(w, "%v", "\n*------------------------------")
}

func index(w http.ResponseWriter, resp *http.Request) {

	message := `🌸🌸 Welcome to the simulator  🌸🌸

*--------------------------------*
| Available Endpoints:           |
| - index     : /                |
| - Dashboard : /dashboard       |
*--------------------------------*

🔑 head to /dashboard that will fetch 
	the responces 

{
  "UID": "<unique identifier>",
  "FirstName": "<user first name>",
  "LastName": "<user last name>",
  "PayAmount": "<user amount>",
  "PaymentMode": "<user payment method>",
  "PayStatus": "<user payment status>",
  "DateOfTransaction": "<user date and time>",
  "Location": "<user location>",
  "Contact": "<user contact number>"
}

🐛 The data will be fetched from the server using go routines
`

	fmt.Fprintf(w, "%v", message)
}

func main(){

	http.HandleFunc("/", index)
	http.HandleFunc("/dashboard", dashboard)
	
	log.Println("\n\nsimulator is running on http://localhost:5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}
