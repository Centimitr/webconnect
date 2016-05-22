package xcache

import (
	msg "github.com/Centimitr/xmessage"
)

type Cache struct{}

func (c Cache) BeforeProcess(ctx *msg.Ctx) {

}

func (c Cache) AfterProcess(ctx *msg.Ctx) {

}

func init() {
	msg.LoadMiddleware(Cache{})
}
