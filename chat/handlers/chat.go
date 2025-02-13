package handlers

import (
	"chat/database"
	"chat/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Active clients list
var clients = make(map[*websocket.Conn]bool)

// Broadcast channel (to send messages to all clients)
var broadcast = make(chan models.Message)

// Websocket upgrader (handle HTTP to websocket)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handle connections
func HandleConnections(c *gin.Context) {
	// Upgrade connection to websocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket: ", err)
		return
	}
	defer ws.Close()

	// Register client
	clients[ws] = true

	// Listen for messages from client
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws) // Remove client if connection is lost
			break
		}

		database.DB.Create(&msg)

		// Send received message to broadcast channel
		broadcast <- msg
	}
}

// Handle messages
func HandleMessages() {
	for {
		msg := <- broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
