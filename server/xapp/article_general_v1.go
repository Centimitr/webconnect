package xapp

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
)

type Article struct{}

func (m Article) GetIndexArticles(ctx *msg.Ctx) {
	ctx.Set(&msg.ParamConfig{Key: "num", Required: false, Default: "100"})
}
func (m Article) GetMessages(ctx *msg.Ctx) {
	fmt.Println("M")
}
func init() {
	var m Article
	msg.LoadModule(m)
}
