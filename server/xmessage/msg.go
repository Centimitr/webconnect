package xmessage

import (
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"reflect"
)

type Msg struct {
	ProcessorMap map[string]*Processor
	// MiddlewareList    map[string][]func()
	AfterReceiveList  []func()
	BeforeProcessList []func()
	AfterProcessList  []func()
	BeforeSendList    []func()
}

/*
	Middleware
*/
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
func (m *Msg) loadMiddleware(x interface{}) {
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

/*
	Processor & Processor Module
*/

type Processor struct {
	Module  string
	PkgPath string
	Name    string
	Func    func(*Ctx) []reflect.Value
}

// Client use a part of the table's key (usually is ProcessorName) to match a processor.
// Currently, the key is "PkgName.ProcessorName" so team members should use different package names to do a replacement though their package path is not the same.
func (m *Msg) matchProcessor(suffix string) (func(*Ctx) []reflect.Value, error) {
	matchedList := []string{}
	// key is "PkgName.ProcessorName"
	for key, _ := range m.ProcessorMap {
		if strings.HasSuffix(key, suffix) {
			matchedList = append(matchedList, key)
		}
	}
	if len(matchedList) == 1 {
		return m.ProcessorMap[matchedList[0]].Func, nil
	}
	return nil, errors.New("Cannot select 1 processor.")
}

// module packages use LoadModule to load itselft
func (m *Msg) loadModule(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	for i := 0; i < v.NumMethod(); i++ {
		index := i
		m.ProcessorMap[t.Name()+"."+t.Method(i).Name] = &Processor{
			PkgPath: t.PkgPath(),
			Module:  t.Name(),
			Name:    t.Method(i).Name,
			Func: func(ctx *Ctx) []reflect.Value {
				return t.Method(index).Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
			},
		}
	}
}

// return a http.Handler
func (m *Middleware) Server(ws *websocket.Conn) {
	var err error
	for {
		var req Req
		if err = websocket.JSON.Receive(ws, &req); err != nil {
			brea
		}
		go do(ws, &req)
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
