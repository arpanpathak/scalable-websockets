package controller

import (
	"net/http"
)

// HandleStaticFile returns websocket client ui
func HandleStaticFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/websocketClient.html")
}
