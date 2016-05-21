package xmessage

import (
	"fmt"
	"reflect"
)

/*
	Middleware
*/
// type Middleware struct {
// 	Name        string
// 	Description string
// }

// type AfterReceiveMiddleware struct{}

// func (mid *AfterReceiveMiddleware) AfterReceive() {
// }

// type BeforeProcessMiddleware struct{}

// func (mid *BeforeProcessMiddleware) BeforeProcess() {
// }

// type AfterProcessMiddleware struct{}

// func (mid *AfterProcessMiddleware) AfterProcess() {
// }

// type BeforeSendMiddleware struct{}

// func (mid *BeforeSendMiddleware) BeforeSend() {
// }

// type AfterSendMiddleware struct{}

// func (mid *AfterSendMiddleware) AfterSend() {
// }

// finally, i think that import relatively middleware package is not a good idea

// type AfterReceiveMiddleware interface {
// 	AfterReceive()
// }

// type BeforeProcessMiddleware interface {
// 	BeforeProcess()
// }

// type AfterProcessMiddleware interface {
// 	AfterProcess()
// }

// type BeforeSendMiddleware interface {
// 	BeforeSend()
// }

// type AfterSendMiddleware interface {
// 	AfterSend()
// }

// middleware packages use LoadMiddleware to load itself
func (m *Msg) loadMiddleware(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Println(t.Name())
	// middlewareName := t.Name()
	// m.Middleware.List[middlewareName].Value = x
	// for i, name := range []string{"AfterReceive", "BeforeProcess", "AfterProcess", "BeforeSend", "AfterSend"} {
	// for i, name := range MIDDLEWARE_STAGE_LIST {
	// if _, ok := t.MethodByName(name); ok {
	// m.Middleware.List[middlewareName].Support[i] = true
	// m.Middleware.Func[i] = append(m.Middleware.Func[i], func() {
	// method.Func.Call(t)
	// })
	// switch name{
	// 	case "AfterReceive"
	// }
}

// }
// switch x := x.(type) {
// case AfterReceiveMiddleware:
// 	m.AfterReceiveList = append(m.AfterReceiveList, x.AfterReceive)
// case BeforeProcessMiddleware:
// 	m.BeforeProcessList = append(m.BeforeProcessList, x.BeforeProcess)
// case AfterProcessMiddleware:
// 	m.AfterProcessList = append(m.AfterProcessList, x.AfterProcess)
// case BeforeSendMiddleware:
// 	m.BeforeSendList = append(m.BeforeSendList, x.BeforeSend)
// case AfterSendMiddleware:
// 	m.AfterSendList = append(m.AfterSendList, x.AfterSend)
// default:
// 	fmt.Println("Error parse type.")
// }
// }
