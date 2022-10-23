package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader for upgrading a http connection to a websocket.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// allowing access to anyone for now.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reader(conn *websocket.Conn) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("!!", err)
		}

		fmt.Println("serverread", string(msg))

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println("@@", err)
			return
		}
	}
}

// websocket function
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrading the http conn to websocket conn using upgrader
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("##", err)
	}

	// Listening to the client for new messages through the websocket
	reader(ws)

}

func addingRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "home route on the server")
	})
	// ws endpoint
	http.HandleFunc("/ws", serveWs)
}

func main() {
	addingRoutes()
	fmt.Println("Chat App")
	http.ListenAndServe(":3002", nil)
}
