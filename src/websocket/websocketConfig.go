package websocket

import "github.com/gorilla/websocket"

// Config WebSocket
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var UpdateSignal = "update"
var Websocket *websocket.Conn
var WsError error
