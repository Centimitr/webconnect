package xmessage

import (
	// "encoding/json"
	"fmt"
	stat "github.com/Centimitr/xmessage/statistics"
	"golang.org/x/net/websocket"
	// "sync"
)

func do(ws *websocket.Conn, req *Req) {
	var err error

	stat.Stat.AddRequest(req.Method)
	// Phase I: initial req,res,ctx
	// err = json.Unmarshal([]byte(reqstr), req)
	res := &Res{Id: req.Id, Method: req.Method}
	ctx := &Ctx{res: res, req: req}
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
	if err := websocket.JSON.Send(ws, res); err != nil {
		fmt.Println("SEND ERROR.")
		return
	}
	stat.Stat.AddResponse(res.Method)
	stat.Stat.Get()
	// resBytes, _ := json.Marshal(*res)
	// if err := websocket.Message.Send(ws, string(resBytes)); err != nil {
	// 	fmt.Println("SEND ERROR.")
	// 	return
	// }
}
