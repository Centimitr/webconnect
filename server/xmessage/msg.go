package xmessage

import (
	"golang.org/x/net/websocket"
)

type middlewareSupport struct {
	Name          string
	AfterReceive  bool
	BeforeProcess bool
	AfterProcess  bool
	BeforeSend    bool
	AfterSend     bool
}

type middleware struct {
	Map               map[string]interface{}
	Support           []middlewareSupport
	AfterReceiveFunc  []func(*Req)
	BeforeProcessFunc []func(*Ctx)
	AfterProcessFunc  []func(*Ctx)
	BeforeSendFunc    []func(*Res)
	AfterSendFunc     []func(*Res)
}

type Msg struct {
	Middleware   middleware
	ProcessorMap map[string]*Processor
}

// return a http.Handler
func (m *Msg) Server(ws *websocket.Conn) {
	var err error
	for {
		var req Req
		if err = websocket.JSON.Receive(ws, &req); err != nil {
			break
		}
		go m.do(ws, &req)
	}
}

/*
	init
*/

var msg *Msg

func init() {
	msg = &Msg{
		Middleware: middleware{
			Map: make(map[string]interface{}),
		},
		ProcessorMap: make(map[string]*Processor),
	}
}

/*
	global
*/

// return msg package's global object: msg
func Ins() *Msg {
	return msg
}

// used in Processor Module package to load itself
func LoadModule(x interface{}) {
	msg.loadModule(x)
}

// used in Middleware package to load itself
func LoadMiddleware(x interface{}) {
	msg.loadMiddleware(x)
}
