package routes

import (
  "fmt"
	"net/http"
)

func NewRoute() http.Handler{

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandlr)
	mux.HandleFunc("/api/data", apiDataHandler)
	
	return mux
}


func indexHandlr(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "test client server..!!")
}


func apiDataHandler(w http.ResponseWriter, r *http.Request){
  date := "ip: 192.168.20.13, subnet: 255.255.255.0 "
	fmt.Fprintln(w, date)
}
