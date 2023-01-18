package controllers

import (
	"github.com/gin-gonic/gin"
	"time"

	mailyWebsocket "maily/go-backend/src/websocket"
)

func WsHandler(c *gin.Context) {
	mailyWebsocket.Websocket, mailyWebsocket.Error = mailyWebsocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if mailyWebsocket.Error != nil {
		return
	}

	// Do not disconnect
	mailyWebsocket.Websocket.SetReadDeadline(time.Time{})
	mailyWebsocket.Websocket.SetWriteDeadline(time.Time{})
}
