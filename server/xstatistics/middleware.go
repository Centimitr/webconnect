package xstatistics

import (
	// "fmt"
	msg "github.com/Centimitr/xmessage"
	"time"
)

/*
	middleware hooked methods
*/

func (s StatisticsMap) AfterReceive(req *msg.Req) {
	req.Temp.Put("Stat", "start", time.Now())
	s.recordReq(req.Method)
}

func (s StatisticsMap) BeforeProcess(ctx *msg.Ctx) {
}

func (s StatisticsMap) AfterProcess(ctx *msg.Ctx) {
}

func (s StatisticsMap) BeforeSend(res *msg.Res) {

}

func (s StatisticsMap) AfterSend(res *msg.Res) {
	t := res.Temp.Get("Stat", "start").(time.Time)
	duration := time.Now().Sub(t)
	s.recordResAndStat(res.Method, duration)
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
