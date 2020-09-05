package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/scalable-websocket/app/controller"
)

func main() {
	// get the appId from the environment variable
	appPORT := fmt.Sprintf(":%s", os.Getenv("APPPORT"))
	appID := fmt.Sprintf(":%s", os.Getenv("APPID"))
	http.HandleFunc("/ws", controller.WebSocketHandler)
	http.HandleFunc("/", controller.HandleStaticFile)

	if err := http.ListenAndServe(appPORT, nil); err != nil {
		panic(err)
	} else {
		fmt.Sprintf("APPID %s is listening on PORT %s", appID, appPORT)
	}
}
