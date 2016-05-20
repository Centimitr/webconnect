package xmessage

import (
	// "encoding/json"
	"fmt"
	stat "github.com/Centimitr/xmessage/statistics"
	"golang.org/x/net/websocket"
	// "sync"
)

func (m *Msg) do(ws *websocket.Conn, req *Req) {
	// var err error

	// Phase I: AfterReceive
	// - initial context and response
	// m.AfterReceive() <- global
	stat.Stat.AddRequest(req.Method)
	res := &Res{Id: req.Id, Method: req.Method}
	ctx := &Ctx{res: res, req: req}
	ctx.Init()

	// Phase II: BeforeProcess
	// - provide chance to manipulate context object for middlewares
	// m.BeforeProcess() <- context

	// Phase III: Process
	// - match and select processor, and then execute it on ctx
	process, _ := m.matchProcessor(req.Method)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	process(ctx)

	// Phase IV: AfterProcess
	// - mainly do response relative tasks
	// m.AfterProcess() <- context
	ctx.setResParams()

	// Phase V: BeforeSend
	// m.BeforeSend() <- global

	// Phase VI: Send
	// - send
	if err := websocket.JSON.Send(ws, res); err != nil {
		fmt.Println("SEND ERROR.")
		return
	}

	// Phase VI: AfterSend
	// m.AfterSend() <- global
	stat.Stat.AddResponse(res.Method)
	stat.Stat.Get()
}
