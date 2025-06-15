package main

import (
	"log"
	"net/http"
	"time"

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

	// Custom HTTP configuration
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panicf("Error starting server: %v", err)
	}
}
