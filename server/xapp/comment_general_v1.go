package xapp

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
)

type Comment struct{}

func (m Comment) GetIndexComments() {
	fmt.Println("C")
}
func (m Comment) GetMessages() {
	fmt.Println("M")
}
func init() {
	var m Comment
	msg.LoadModule(m)
}
