package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Greeting string `json:"greeting"`
}

func main() {
	// Force log's color
	gin.ForceConsoleColor()

	// Create a new Gin router
	router := gin.Default()

	// Define a handler function for the "/hello" endpoint
	router.GET("/hello", func(c *gin.Context) {
		greeting := Message{
			Greeting: "Hello, World!",
		}
		c.JSON(http.StatusOK, greeting)
	})

	// Start the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		log.Panicf("Error starting server: %v", err)
	}
}
