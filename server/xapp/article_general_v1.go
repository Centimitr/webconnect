package xapp

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
)

type Article struct{}

func (m Article) GetIndexArticles(req *msg.Req, res *msg.Res) {
	fmt.Println(req.Id, res.Id)
}
func (m Article) GetMessages(req *msg.Req, res *msg.Res) {
	fmt.Println("M")
}
func init() {
	var m Article
	msg.LoadModule(m)
}
