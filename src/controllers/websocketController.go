package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	mailyWebsocket "maily/go-backend/src/websocket"
)

func WsHandler(c *gin.Context) {
	mailyWebsocket.Websocket, mailyWebsocket.Error = mailyWebsocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if mailyWebsocket.Error != nil {
		fmt.Println(mailyWebsocket.Error)
		return
	}

	// Do not disconnect
	mailyWebsocket.Websocket.SetReadDeadline(time.Time{})
	mailyWebsocket.Websocket.SetWriteDeadline(time.Time{})
}
