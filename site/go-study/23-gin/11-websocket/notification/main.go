package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		for {
			err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
