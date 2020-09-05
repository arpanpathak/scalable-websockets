package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

// Upgrader upgrade header
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// WebSocketHandler handle websocket requests
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("New upgrade hader request.........")

	// Upgrade HTTP 1.1 request to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	// Panic the error
	if err != nil {
		panic(err)
	}

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n from %s", conn.RemoteAddr(), string(msg), os.Getenv("APPID"))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
