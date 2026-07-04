package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)





var upgrader = websocket.Upgrader{
	CheckOrigin : func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn,err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		return 
	}
	defer conn.Close()

	for {
		_,msg,err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("Message Hello.", string(msg))
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}
func main() {
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe(":8080",nil)
}