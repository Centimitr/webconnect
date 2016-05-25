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
	// s.timesStatAR(req.Method)
	// s.durationStatAR(req.Id)
	// req.Temp["Statistic"]["start"] = time.Now().String()
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
	// fmt.Println(res.Temp["Statistics"]["start"])
	t := res.Temp.Get("Stat", "start").(time.Time)
	duration := time.Now().Sub(t)
	s.recordResAndStat(res.Method, duration)
	// s.timesStatAS(res.Method)
	// s.durationStatAS(res.Method, res.Id)
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
			// timeMap:   make(map[string]time.Time),
		},
	})
}
