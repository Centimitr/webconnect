package xapp

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
)

type Article struct{}

func (m Article) GetIndexArticles(c *msg.Ctx) {
	c.Set(&msg.ParamConfig{Key: "num", Required: false, Default: "100", Echo: true})
	c.Set(&msg.ParamConfig{Key: "a", Required: false, Echo: true})
	if c.Error.Fatal == nil {
		fmt.Println("No Fatal Error.")
		p1 := c.Get("num")
		p2 := c.Get("a")
		c.Params["a"] = `456"""7"8""9`
		fmt.Println(p1, p2)
		fmt.Println(c.Params)
	}
}
func (m Article) GetMessages(ctx *msg.Ctx) {
	fmt.Println("M")
}
func init() {
	var m Article
	msg.LoadModule(&m)
}
