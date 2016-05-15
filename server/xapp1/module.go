package xapp1

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
)

func init() {
	var m Module
	msg.Load(m)
}

type Module struct{}

func (m Module) GetIndexArticles() {
	fmt.Println("A")
}
func (m Module) GetMessages() {
	fmt.Println("M")
}
