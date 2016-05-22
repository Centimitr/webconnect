package xmessage

import (
	"fmt"
	"reflect"
)

// middleware packages use LoadMiddleware to load itself
func (m *Msg) loadMiddleware(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	middlewareName := t.Name()
	m.Middleware.Map[middlewareName] = x
	var support = middlewareSupport{
		Name: middlewareName,
	}
	for _, name := range []string{"AfterReceive", "BeforeProcess", "AfterProcess", "BeforeSend", "AfterSend"} {
		if method, ok := t.MethodByName(name); ok {
			switch name {
			case "AfterReceive":
				support.AfterReceive = true
				m.Middleware.AfterReceiveFunc = append(m.Middleware.AfterReceiveFunc, func(req *Req) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(req)})
				})
			case "BeforeProcess":
				support.BeforeProcess = true
				m.Middleware.BeforeProcessFunc = append(m.Middleware.BeforeProcessFunc, func(ctx *Ctx) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
				})
			case "AfterProcess":
				support.AfterProcess = true
				m.Middleware.AfterProcessFunc = append(m.Middleware.AfterProcessFunc, func(ctx *Ctx) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
				})
			case "BeforeSend":
				support.BeforeSend = true
				m.Middleware.BeforeSendFunc = append(m.Middleware.BeforeSendFunc, func(res *Res) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(res)})
				})
			case "AfterSend":
				support.AfterSend = true
				m.Middleware.AfterSendFunc = append(m.Middleware.AfterSendFunc, func(res *Res) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(res)})
				})
			default:
				fmt.Println("Middleware load logic error.")
			}
		}
	}
	m.Middleware.Support = append(m.Middleware.Support, support)
}

/*
	exec methods
*/

func (m *Msg) AfterReceive(req *Req) {
	for _, fn := range m.Middleware.AfterReceiveFunc {
		fn(req)
	}
}

func (m *Msg) BeforeProcess(ctx *Ctx) {
	for _, fn := range m.Middleware.BeforeProcessFunc {
		fn(ctx)
	}
}

func (m *Msg) AfterProcess(ctx *Ctx) {
	for _, fn := range m.Middleware.AfterProcessFunc {
		fn(ctx)
	}
}

func (m *Msg) BeforeSend(res *Res) {
	for _, fn := range m.Middleware.BeforeSendFunc {
		fn(res)
	}
}

func (m *Msg) AfterSend(res *Res) {
	for _, fn := range m.Middleware.AfterSendFunc {
		fn(res)
	}
}
