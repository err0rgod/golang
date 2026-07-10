package main

import (
	"net/http"
	"fmt"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin : func (r *http.Request) bool {
		return true
	},
}

type Data struct {
	Type int
	Data []byte
}

var Register  = make(chan *websocket.Conn)
var Unregister  = make(chan *websocket.Conn)
var broadcast = make(chan Data)

func hub() {
	Clients := make(map[*websocket.Conn]bool)
	for {
		select {
		case conn := <- Register:
			Clients[conn]=true
		
		case conn := <- Unregister:
			delete(Clients,conn)
			conn.Close()

		case message := <- broadcast:
			for client := range Clients {
				client.WriteMessage(message.Type,message.Data)
				// if err != nil  {return}
			}
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// adding clients
	Register <- conn
	defer func(){
		Unregister <- conn
	}()

	// just reading messages 
	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil{
			return 
		}
		broadcast<- Data{
			Type : msgType,
			Data : message ,	
		}
	}

	
}

func WebsocketHandler() {
	go hub()
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("Server chal raha hai :8080")
	http.ListenAndServe(":8080", nil)
}
