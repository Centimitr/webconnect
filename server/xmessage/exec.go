package xmessage

import (
	// "encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	// "sync"
)

// var Stat *StatisticsMap

// func init() {
// 	Stat = &StatisticsMap{
// 		m: make(map[string]*Statistics),
// 	}
// }

// type Statistics struct {
// 	// Method        string
// 	RequestTimes  int
// 	ResponseTimes int
// }

// type StatisticsMap struct {
// 	lock sync.RWMutex
// 	m    map[string]*Statistics
// }

// func (s *StatisticsMap) AddRequest(method string) {
// 	// s.lock.Lock()
// 	// defer s.lock.Unlock()
// 	s.m[method].RequestTimes++
// }

// func (s *StatisticsMap) AddResponse(method string) {
// 	// s.lock.Lock()
// 	// defer s.lock.Unlock()
// 	s.m[method].ResponseTimes++
// }

// func (s *StatisticsMap) Get() {
// 	// s.lock.Lock()
// 	// defer s.lock.Unlock()
// 	for k, item := range s.m {
// 		fmt.Println(k, item.RequestTimes, item.ResponseTimes)
// 	}
// }

func do(ws *websocket.Conn, req *Req) {
	var err error

	// Stat.AddRequest(req.Method)
	// Phase I: initial req,res,ctx
	// err = json.Unmarshal([]byte(reqstr), req)
	res := &Res{Id: req.Id, Method: req.Method}
	ctx := &Ctx{res: res, req: req}
	ctx.Init()
	// Phase II: context processs, match -> process
	process, err := matchProcessor(req.Method)
	if err != nil {
		fmt.Println(err)
		return
	}
	process(ctx)

	// Phase III: response setting
	ctx.setResParams()

	// Phase IV: response call
	// websocket response
	if err := websocket.JSON.Send(ws, res); err != nil {
		fmt.Println("SEND ERROR.")
		return
	}
	// Stat.AddResponse(res.Method)
	// Stat.Get()
	// resBytes, _ := json.Marshal(*res)
	// if err := websocket.Message.Send(ws, string(resBytes)); err != nil {
	// 	fmt.Println("SEND ERROR.")
	// 	return
	// }
}
