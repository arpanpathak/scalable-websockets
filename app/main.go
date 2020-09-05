package main

import (
	"net/http"
	"os"

	"github.com/scalable-websocket/app/controller"
)

func main() {
	// get the appId from the environment variable
	appPORT := os.Getenv("APPPORT")

	http.HandleFunc("/ws", controller.WebSocketHandler)

	http.HandleFunc("/", controller.HandleStaticFile)

	if err := http.ListenAndServe(appPORT, nil); err != nil {
		panic(err)
	}
}
