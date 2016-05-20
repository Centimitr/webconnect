package xmessage

import (
	"fmt"
)

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

type AfterSendMiddleware struct{}

func (mid *AfterSendMiddleware) AfterSend() {
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
	case AfterSendMiddleware:
		m.AfterSendList = append(m.AfterSendList, x.AfterSend)
	default:
		fmt.Println("Error parse type.")
	}
}
