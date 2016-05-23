package xjsonbase

import (
	msg "github.com/Centimitr/xmessage"
)

/*
	middleware hooked methods
*/

// func (j JSONBase) AfterReceive(req *msg.Req) {
// }

// func (j JSONBase) BeforeProcess(ctx *msg.Ctx) {
// }

// func (j JSONBase) AfterProcess(ctx *msg.Ctx) {
// }

// func (j JSONBase) BeforeSend(res *msg.Res) {

// }

// func (j JSONBase) AfterSend(res *msg.Res) {
// }

/*
	init
*/

func init() {
	msg.LoadMiddleware(JSONBase{})
}
