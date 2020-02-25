package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Message string `json:"message"`
}

func RunWs() {
	http.HandleFunc("/ws", handleConnections)
	err := http.ListenAndServe(":8042", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer ws.Close()
	clients[ws] = true
}

func SendMsg(msg string) {
	message := Message{
		Message: msg,
	}
	for client := range clients {
		err := client.WriteJSON(&message)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
