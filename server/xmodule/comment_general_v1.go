package xmodule

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
)

type Comment struct{}

func (m Comment) GetIndexComments(ctx *msg.Ctx) {
	fmt.Println("C")
}
func (m Comment) GetMessages(ctx *msg.Ctx) {
	fmt.Println("M")
}
func init() {
	var m Comment
	msg.LoadModule(m)
}
