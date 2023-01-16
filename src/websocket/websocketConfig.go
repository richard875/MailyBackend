package websocket

import "github.com/gorilla/websocket"

// Config WebSocket
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var UpdateSignal = "update"
var Delimiter = "##__##"
var Websocket *websocket.Conn
var Error error
