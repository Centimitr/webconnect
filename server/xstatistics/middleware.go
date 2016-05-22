package xstatistics

import (
	msg "github.com/Centimitr/xmessage"
)

/*
	middleware hooked methods
*/

func (s StatisticsMap) AfterReceive(req *msg.Req) {
	s.addRequest(req.Method)
}

func (s StatisticsMap) BeforeProcess(ctx *msg.Ctx) {

}

func (s StatisticsMap) AfterProcess(ctx *msg.Ctx) {

}

func (s StatisticsMap) BeforeSend(res *msg.Res) {

}

func (s StatisticsMap) AfterSend(res *msg.Res) {
	s.addResponse(res.Method)
	s.get()
}

/*
	init
*/

type Statistcs struct {
	StatisticsMap
}

func init() {
	msg.LoadMiddleware(Statistcs{
		StatisticsMap{
			methodMap: make(map[string]*StatisticsItem),
		},
	})
}
