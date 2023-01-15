package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	mailyWebsocket "maily/go-backend/src/websocket"
)

func WsHandler(c *gin.Context) {
	mailyWebsocket.Websocket, mailyWebsocket.WsError = mailyWebsocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if mailyWebsocket.WsError != nil {
		fmt.Println(mailyWebsocket.WsError)
		return
	}

	// Disconnect after 5 minutes
	mailyWebsocket.Websocket.SetReadDeadline(time.Now().Add(300 * time.Second))
	mailyWebsocket.Websocket.SetWriteDeadline(time.Now().Add(300 * time.Second))
}
