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
}

/*
	init
*/

// use struct Statistics to combime the middleware, so Statistics is the middleware name.
type Statistics struct {
	StatisticsMap
}

func init() {
	msg.LoadMiddleware(Statistics{
		StatisticsMap{
			methodMap: make(map[string]*StatisticsItem),
		},
	})
}
