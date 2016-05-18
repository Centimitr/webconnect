package xmessage

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

func do(ws *websocket.Conn, reqstr string) {
	var err error

	// Phase I: initial req,res,ctx
	var req *Req
	var res *Res
	var ctx *Ctx
	err = json.Unmarshal([]byte(reqstr), req)
	res = &Res{Id: req.Id}
	ctx = &Ctx{res: res, req: req}
	ctx.Init()
	// Phase II: context processs, match -> process
	process, err := matchProcessor(req.Method)
	if err != nil {
		fmt.Println(err)
		return
	}
	process(ctx)

	// Phase III: response setting
	ctx.setResParams()

	// Phase IV: response call
	// websocket response
	resBytes, _ := json.Marshal(*res)
	if err := websocket.Message.Send(ws, string(resBytes)); err != nil {
		fmt.Println("SEND ERROR.")
		return
	}
}
