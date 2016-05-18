package xmessage

import (
	"golang.org/x/net/websocket"
)

func Server(ws *websocket.Conn) {
	var err error
	for {
		var req Req
		if err = websocket.JSON.Receive(ws, &req); err != nil {
			break
		}
		go do(ws, &req)
	}
}
