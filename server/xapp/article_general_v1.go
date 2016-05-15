package xapp

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
)

type Article struct{}

func (m Article) GetIndexArticles() {
	fmt.Println("A")
}
func (m Article) GetMessages() {
	fmt.Println("M")
}
func init() {
	var m Article
	msg.LoadModule(m)
}
