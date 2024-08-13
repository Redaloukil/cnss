package main

import (
	"fmt"
	"net/http"
	"github.com/Himanshu-Parangat/cnss/internal/routes"
)

func main(){
	router := routes.NewRoute() 

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening an http://localhost%s\n", addr)

	err := http.ListenAndServe(addr, router)

	if err != nil{
		panic(err)
	}

}
