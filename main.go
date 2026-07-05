package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var broadcast = make(chan []byte)	

var Clients = make(map[*websocket.Conn]bool)

func hub() {
	for {
		message := <- broadcast
		for client := range Clients {
			client.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// adding clients
	Clients[conn] = true
	defer func(){
		delete(Clients, conn)
		conn.Close()
	}()

	// just reading messages 
	for {
		_, message, err := conn.ReadMessage()
		if err != nil{
			return 
		}
		broadcast<- message
	}

	
}

func main() {
	go hub()
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("Server chal raha hai :8080")
	http.ListenAndServe(":8080", nil)
}