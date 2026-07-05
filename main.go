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

type Message struct {
	Type int
	Data []byte
}


var register = make(chan *websocket.Conn)
var unregister = make(chan *websocket.Conn)
var broadcast = make(chan Message)

// only writes the messages 
func hub() {
	Clients := make(map[*websocket.Conn]bool)
	for {
		select {
		case conn := <- register:
			Clients[conn]=true
		
		case conn := <- unregister:
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
	register <- conn
	defer func(){
		unregister <- conn
	}()

	// just reading messages 
	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil{
			return 
		}
		broadcast<- Message{
			Type : msgType,
			Data : message ,	
		}
	}

	
}

func main() {
	go hub()
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("Server chal raha hai :8080")
	http.ListenAndServe(":8080", nil)
}