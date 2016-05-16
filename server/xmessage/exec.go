package xmessage

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

func do(ws *websocket.Conn, req *Req) {
	// format req
	// match
	f, err := matchProcessor(req.Method)
	if err != nil {
		fmt.Println(err)
		return
	}
	res := &Res{Id: req.Id}
	f(req, res)
	// send back
	resBytes, _ := json.Marshal(*res)
	if err := websocket.Message.Send(ws, string(resBytes)); err != nil {
		fmt.Println("SEND ERROR.")
		return
	}
}
