package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func message_send(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {

	fmt.Println("\n\nclient running..")
	router := gin.Default()

	router.GET("/ping", message_send)
	router.Run()
}
