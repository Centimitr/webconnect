package xstatistics

import (
	msg "github.com/Centimitr/xmessage"
	"time"
)

/*
	middleware hooked methods
*/

func (s StatisticsMap) AfterReceive(req *msg.Req) {
	s.addRequest(req.Method)
	s.timeMap[req.Id] = time.Now()
}

func (s StatisticsMap) BeforeProcess(ctx *msg.Ctx) {
}

func (s StatisticsMap) AfterProcess(ctx *msg.Ctx) {
}

func (s StatisticsMap) BeforeSend(res *msg.Res) {

}

func (s StatisticsMap) AfterSend(res *msg.Res) {
	s.addResponse(res.Method)
	s.lock.Lock()
	defer s.lock.Unlock()
	duration := time.Now().Sub(s.timeMap[res.Id]).Seconds()
	if s.methodMap[res.Method].MinDuration < 1e-6 {
		s.methodMap[res.Method].MinDuration = duration
	}
	switch {
	case s.methodMap[res.Method].MinDuration-duration > 1e-6:
		s.methodMap[res.Method].MinDuration = duration
	case duration-s.methodMap[res.Method].MaxDuration > 1e-6:
		s.methodMap[res.Method].MaxDuration = duration
	case true:
		times := s.methodMap[res.Method].ResponseTimes
		s.methodMap[res.Method].AvgDuration = (float64(times-1)*s.methodMap[res.Method].AvgDuration + duration) / float64(times)
	}
	delete(s.timeMap, res.Id)
	if len(s.timeMap) > 64000 {
		for id, t := range s.timeMap {
			d := time.Now().Sub(t)
			if d > time.Second {
				delete(s.timeMap, id)
			}
		}
	}
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
			timeMap:   make(map[string]time.Time),
		},
	})
}
