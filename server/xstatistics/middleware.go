package xstatistics

import (
	msg "github.com/Centimitr/xmessage"
)

/*
	middleware hooked methods
*/

func (s StatisticsMap) AfterReceive(req *msg.Req) {
	s.addRequest("req")
}

func (s StatisticsMap) BeforeProcess(ctx *msg.Ctx) {

}

func (s StatisticsMap) AfterProcess(ctx *msg.Ctx) {

}

func (s StatisticsMap) BeforeSend(res *msg.Res) {

}

func (s StatisticsMap) AfterSend() {
	s.addResponse("res")
	s.get()
}
