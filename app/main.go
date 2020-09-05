package main

import (
	"net/http"

	"github.com/scalable-websocket/app/controller"
)

func main() {
	// get the appId from the environment variable
	appPort := ":8081"

	http.HandleFunc("/echo", controller.WebSocketHandler)

	http.HandleFunc("/", controller.HandleStaticFile)

	if err := http.ListenAndServe(appPort, nil); err != nil {
		panic(err)
	}
}
