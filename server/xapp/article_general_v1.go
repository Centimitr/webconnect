package xapp

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
)

type Article struct{}

func (m Article) GetIndexArticles(c *msg.Ctx) {
	c.Set(&msg.ParamConfig{Key: "num", Required: false, Default: "100"})
	c.Set(&msg.ParamConfig{Key: "a", Required: false})
	if c.Error.Fatal == nil {
		fmt.Println("No Fatal Error.")
		p1 := c.Get("num")
		p2 := c.Get("a")
		fmt.Println(p1, p2)
		fmt.Println(c.Params)
	}
}
func (m Article) GetMessages(ctx *msg.Ctx) {
	fmt.Println("M")
}
func init() {
	var m Article
	msg.LoadModule(m)
}
