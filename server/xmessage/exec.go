package xmessage

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

func do(ws *websocket.Conn, req *Req) {
	// format req
	res := &Res{Id: req.Id}
	ctx := &Ctx{res: res, req: req} // use req
	ctx.Init()
	// match
	f, err := matchProcessor(req.Method)
	if err != nil {
		fmt.Println(err)
		return
	}
	f(ctx)
	// send back
	resBytes, _ := json.Marshal(*res)
	if err := websocket.Message.Send(ws, string(resBytes)); err != nil {
		fmt.Println("SEND ERROR.")
		return
	}
}
