package handlers

import (
	"fmt"
	"net/http"

	"github.com/adamlahbib/go-realtimechat-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan models.Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(c *gin.Context) {
	w := c.Writer
	r := c.Request

	UserEmail := c.Query("useremail")
	RommName := c.Query("roomname")

	fmt.Println(UserEmail)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if UserEmail != "" {
		clients[conn] = UserEmail
	} else {
		clients[conn] = RommName
	}

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			return
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client, clientUsername := range clients {
			fmt.Println("Sending message to", clientUsername)
			fmt.Println(msg.Receiver)
			if clientUsername == msg.Receiver {
				err := client.WriteJSON(msg)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
