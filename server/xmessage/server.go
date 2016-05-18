package xmessage

import (
	"golang.org/x/net/websocket"
)

func Server(ws *websocket.Conn) {
	var err error
	for {
		var req string
		if err = websocket.Message.Receive(ws, &req); err != nil {
			break
		}
		go do(ws, req)
	}
}
