package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func postUser(r *gin.Engine) *gin.Engine {
	r.POST("/user/add", func(c *gin.Context) {
		var user User
		err := c.BindJSON(&user)
		if err != nil {
			log.Fatalf("Failed to bind JSON: %v\n", err)
		}
		c.JSON(200, user)
	})
	return r
}

func main() {
	r := setupRouter()
	r = postUser(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
