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
	data := process_json_data(get_data())
	fmt.Fprintf(w, "<h1>dashboard<h1></br>")
	fmt.Fprintf(w, `
		<html>
		<head>
			<title>Dashboard</title>
			<style>
				body { font-family: Arial, sans-serif; padding: 20px; }
				table { width: 100%%; border-collapse: collapse; margin-top: 20px; }
				th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
				th { background-color: #f4f4f4; }
			</style>
		</head>
		<body>
			<h1>Transaction Dashboard</h1>
			<table>
				<tr>
					<th>UID</th>
					<th>First Name</th>
					<th>Last Name</th>
					<th>Pay Amount</th>
					<th>Payment Mode</th>
					<th>Pay Status</th>
					<th>Date</th>
					<th>Location</th>
					<th>Contact</th>
				</tr>
				<tr>
					<td>%s</td>
					<td>%s</td>
					<td>%s</td>
					<td>$%v</td>
					<td>%s</td>
					<td>%s</td>
					<td>%s</td>
					<td>%s</td>
					<td>%s</td>
				</tr>
			</table>
		</body>
		</html>
	`,
		data["UID"], data["FirstName"], data["LastName"],
		data["PayAmount"], data["PaymentMode"], data["PayStatus"],
		data["DateOfTransaction"], data["Location"], data["Contact"],
	)
}


func main() {
	fmt.Println("\n\nclient running..")

	http.HandleFunc("/dashboard", dashboard)
	http.ListenAndServe(":8080", nil)
}
