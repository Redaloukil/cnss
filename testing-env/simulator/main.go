package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetServerAddress()  string{
	host := os.Getenv("SERVER_HOST_IP")
	port := os.Getenv("SERVER_PORT")

	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf("%s:%s", host, port)
}


func GetSimulatorAddress() string {
	host := os.Getenv("SIMULATOR_HOST_IP")
	port := os.Getenv("SIMULATOR_PORT")

	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "7000"
	}

	return fmt.Sprintf("%s:%s", host, port)
}

func get_data() map[string]interface{} {

	resp, err := http.Get("http://" + GetServerAddress() + "/api/v1/data")
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
	
	address := GetSimulatorAddress()
	log.Println("\n\nsimulator is running on http://"+address)
	log.Fatal(http.ListenAndServe(address, nil))
}
