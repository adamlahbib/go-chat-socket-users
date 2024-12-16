package handlers

import (
	"fmt"

	"github.com/adamlahbib/go-realtimechat-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var RoomClients = make(map[*websocket.Conn]string)
var RoomBroadcast = make(chan models.Message)

func HandleRoomConnections(c *gin.Context) {
	w := c.Writer
	r := c.Request

	RoomName := c.Query("roomname")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = RoomName

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			break
		}
		broadcast <- msg
	}
}

func HandleRoomMessages() {
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
