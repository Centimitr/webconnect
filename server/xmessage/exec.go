package xmessage

import (
	"fmt"
	// stat "github.com/Centimitr/xmessage/statistics"
	"encoding/json"
	"golang.org/x/net/websocket"
)

func (m *Msg) do(ws *websocket.Conn, req *Req) {
	var err error

	// Phase I: AfterReceive
	// - global, req relative methods
	m.AfterReceive(req)
	// - initial context and response
	res := &Res{Id: req.Id, Method: req.Method}
	ctx := &Ctx{res: res, req: req, Method: req.Method, Data: req.Data}
	ctx.Init()

	// Phase II: BeforeProcess
	// - global, ctx relative methods
	// - provide chance to manipulate context object for middlewares
	m.BeforeProcess(ctx)

	// Phase III: Process
	// - match and select processor, and then execute it on ctx
	process, err := m.matchProcessor(req.Method)
	if err != nil {
		fmt.Println(err)
		return
	}
	process(ctx)

	// Phase IV: AfterProcess
	// - global, ctx relative methods
	m.AfterProcess(ctx)
	// - mainly do response relative tasks
	// res.Data = ctx.Data
	data, err := json.Marshal(ctx.Data)
	if err != nil {
		fmt.Println(err)
	}
	res.Data = string(data)
	ctx.setResParams()

	// Phase V: BeforeSend
	m.BeforeSend(res)

	// Phase VI: Send
	// - send
	if err := websocket.JSON.Send(ws, res); err != nil {
		fmt.Println("SEND ERROR.", err)
		return
	}

	// Phase VI: AfterSend
	m.AfterSend(res)
}
