package xmessage

import (
	"golang.org/x/net/websocket"
)

type Msg struct {
	ProcessorMap map[string]*Processor
	// MiddlewareList    map[string][]func()
	AfterReceiveList  []func()
	BeforeProcessList []func()
	AfterProcessList  []func()
	BeforeSendList    []func()
	AfterSendList     []func()
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
