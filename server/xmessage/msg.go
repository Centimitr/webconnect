package xmessage

import (
	"golang.org/x/net/websocket"
)

var MIDDLEWARE_STAGE_LIST = []string{"AfterReceive", "BeforeProcess", "AfterProcess", "BeforeSend", "AfterSend"}

type middlewareList map[string]middlewareListItem

type middlewareListItem struct {
	Value interface{}
	// Support struct {
	// 	AfterReceive  bool
	// 	BeforeProcess bool
	// 	AfterProcess  bool
	// 	BeforeSend    bool
	// 	AfterSend     bool
	// }
}

type Msg struct {
	// AfterReceiveList  []func()
	// BeforeProcessList []func()
	// AfterProcessList  []func()
	// BeforeSendList    []func()
	// AfterSendList     []func()
	Middleware struct {
		List middlewareList
		// Func [5]reflect.
		// AfterReceiveFunc  []func()
		// BeforeProcessFunc []func()
		// AfterProcessFunc  []func()
		// BeforeSendFunc    []func()
		// AfterSendFunc     []func()
	}
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
		Middleware: struct {
			List map[string]middlewareListItem
			// Func [5]reflect.
			// AfterReceiveFunc  []func()
			// BeforeProcessFunc []func()
			// AfterProcessFunc  []func()
			// BeforeSendFunc    []func()
			// AfterSendFunc     []func()
		}{
			List: make(map[string]middlewareListItem),
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
