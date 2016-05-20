package xmessage

import (
	"fmt"
)

type Msg struct {
	ProcessorMap map[string]*Processor
	// MiddlewareList    map[string][]func()
	AfterReceiveList  []func()
	BeforeProcessList []func()
	AfterProcessList  []func()
	BeforeSendList    []func()
}

type Middleware struct {
	Name        string
	Description string
}

type AfterReceiveMiddleware struct{}

func (mid *AfterReceiveMiddleware) AfterReceive() {
}

type BeforeProcessMiddleware struct{}

func (mid *BeforeProcessMiddleware) BeforeProcess() {
}

type AfterProcessMiddleware struct{}

func (mid *AfterProcessMiddleware) AfterProcess() {
}

type BeforeSendMiddleware struct{}

func (mid *BeforeSendMiddleware) BeforeSend() {
}

// middleware packages use LoadMiddleware to load itself
func (m *Msg) LoadMiddleware(x interface{}) {
	switch x := x.(type) {
	case AfterReceiveMiddleware:
		m.AfterReceiveList = append(m.AfterReceiveList, x.AfterReceive)
	case BeforeProcessMiddleware:
		m.BeforeProcessList = append(m.BeforeProcessList, x.BeforeProcess)
	case AfterProcessMiddleware:
		m.AfterProcessList = append(m.AfterProcessList, x.AfterProcess)
	case BeforeSendMiddleware:
		m.BeforeSendList = append(m.BeforeSendList, x.BeforeSend)
	default:
		fmt.Println("Error parse type.")
	}
}

func (m *Msg) LoadModule(p *Processor) {
	m.ProcessorMap[p.Module+"."+p.Name] = p
}

// func (m *Middleware) Server(ws *websocket.Conn) {
// 	var err error
// 	for {
// 		var req Req
// 		if err = websocket.JSON.Receive(ws, &req); err != nil {
// 			brea
// 		}
// 		go do(ws, &req)
// 	}
// }

/*
	init
*/

var msg *Msg

func init() {
	msg = &Msg{
		ProcessorMap: make(map[string]*Processor),
	}
}

func Ins() *Msg {
	return msg
}
