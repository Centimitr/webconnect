package xcache

import (
	msg "github.com/Centimitr/xmessage"
)

type Cache struct{}

func (c Cache) BeforeProcess() {

}

func (c Cache) AfterProcess() {

}

func init() {
	msg.LoadMiddleware(Cache{})
}
