package xmessage

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

func Server(ws *websocket.Conn) {
	var err error
	for {
		var req string
		if err = websocket.Message.Receive(ws, &req); err != nil {
			fmt.Println("WAITING.")
			break
		}
		var r Req
		json.Unmarshal([]byte(req), &r)
		go do(ws, &r)
	}
}
